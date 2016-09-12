package main

import (
	"bufio"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8081", "http service address")

var abc = `
{"action": "LOGIN", "username": "linwe"}
{"action": "CREATE_ROOM"}
{"action": "JOIN_ROOM", "room_id": 1}
`

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	log.Printf("connected to %s", u.String())
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	for {
		bio := bufio.NewReader(os.Stdin)
		inputStr, _, err := bio.ReadLine()
		if len(inputStr) > 0 {
			err = c.WriteMessage(websocket.TextMessage, []byte(inputStr))
			if err != nil {
				log.Println("Write failed:", err)
				return
			}
		}
	}
}
