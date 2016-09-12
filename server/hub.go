package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is an middleman between the websocket connection and the hub.
type Connection struct {
	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	sendBuffer chan []byte
}

type ConnectionStateChangeListener interface {
	notifyRegister(*Connection) error
	notifyUnregister(*Connection) error
}

type IncomingMessageListener interface {
	notifyIncomingMessage(*Connection, []byte) error
}

type Hub struct {
	connections map[*Connection]bool

	register   chan *Connection
	unregister chan *Connection

	connectionStateChangeListeners []ConnectionStateChangeListener
	incomingMessageListeners       []IncomingMessageListener
}

// readPump pumps messages from the websocket connection to the hub.
func (c *Connection) readPump(hub *Hub) {
	defer func() {
		hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}

		for _, listener := range hub.incomingMessageListeners {
			listener.notifyIncomingMessage(c, message)
		}
	}
}

// write writes a message with the given message type and payload.
func (c *Connection) write(mt int, payload []byte) error {
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.conn.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (c *Connection) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.sendBuffer:
			if !ok {
				// The hub closed the channel.
				c.write(websocket.CloseMessage, []byte{})
				return
			}

			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.sendBuffer)
			for i := 0; i < n; i++ {
				w.Write(<-c.sendBuffer)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func NewHub() *Hub {
	return &Hub{
		connections: make(map[*Connection]bool),
		register:    make(chan *Connection),
		unregister:  make(chan *Connection),
	}
}

func (h *Hub) appendConnectionStateChangeLIstener(
	listener ConnectionStateChangeListener) {
	h.connectionStateChangeListeners =
		append(h.connectionStateChangeListeners, listener)
}

func (h *Hub) appendIncomingMessageLIstener(
	listener IncomingMessageListener) {
	h.incomingMessageListeners =
		append(h.incomingMessageListeners, listener)
}

func (h *Hub) loop() {
	for {
		select {
		case connection := <-h.register:
			h.connections[connection] = true
			for _, listener := range h.connectionStateChangeListeners {
				listener.notifyRegister(connection)
			}
		case connection := <-h.unregister:
			if _, ok := h.connections[connection]; ok {
				for _, listener := range h.connectionStateChangeListeners {
					listener.notifyUnregister(connection)
				}
				delete(h.connections, connection)
				close(connection.sendBuffer)
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func (h *Hub) HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Connection{
		conn:       conn,
		sendBuffer: make(chan []byte, 25600),
	}
	h.register <- client
	go client.writePump()
	client.readPump(h)
}
