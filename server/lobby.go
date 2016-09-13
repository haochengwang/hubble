package main

import (
	"encoding/json"
	"log"
)

const (
	ACTIVE int8 = iota
	AFK
	LOST_CONN
)

type UserMessageAction int64
type UserId int64
type RoomId int64

const (
	LOGIN UserMessageAction = iota
	CREATE_ROOM
	JOIN_ROOM
	LEAVE_ROOM

	STATE_CHANGE
)

type User struct {
	id       UserId
	username string
	nickname string
	conn     *Connection
	state    int8
}

type Lobby struct {
	// Users and connections
	anonymousConn map[*Connection]bool
	userConnMap   map[*Connection]*User
	userStateMap  map[int8]map[UserId]*User

	// Rooms
	rooms         map[RoomId]*Room
	userInRoomMap map[UserId]*Room

	// Utils
	idGen *IdGenerator
}

func NewLobby() *Lobby {
	userStateMap := make(map[int8]map[UserId]*User)
	userStateMap[ACTIVE] = make(map[UserId]*User)
	userStateMap[AFK] = make(map[UserId]*User)
	userStateMap[LOST_CONN] = make(map[UserId]*User)
	return &Lobby{
		anonymousConn: make(map[*Connection]bool),
		userConnMap:   make(map[*Connection]*User),
		userStateMap:  userStateMap,
		rooms:         make(map[RoomId]*Room),
		userInRoomMap: make(map[UserId]*Room),
		idGen:         NewIdGenerator(1),
	}
}

func (l *Lobby) notifyRegister(conn *Connection) (err error) {
	l.anonymousConn[conn] = true
	return nil
}

func (l *Lobby) notifyUnregister(conn *Connection) (err error) {
	if _, ok := l.anonymousConn[conn]; ok {
		delete(l.anonymousConn, conn)
	} else if user, ok := l.userConnMap[conn]; ok {
		delete(l.userConnMap, conn)
		delete(l.userStateMap[user.state], user.id)
		l.userStateMap[LOST_CONN][user.id] = user

		// TODO: schedule LOST_CONN event
		user.state = LOST_CONN
	} else {
		// Should never happen
	}
	return nil
}

func (l *Lobby) notifyIncomingMessage(conn *Connection, message []byte) (err error) {
	log.Println("notifyIncomingMessage: %v", string(message))
	var userMessage map[string]interface{}
	if err = json.Unmarshal(message, &userMessage); err != nil {
		return err
	}
	if _, ok := l.anonymousConn[conn]; ok {
		action := userMessage["action"].(string)
		switch action {
		case "LOGIN":
			return l.handleLoginMessage(conn, userMessage)
		}
	} else if _, ok := l.userConnMap[conn]; ok {
		action := userMessage["action"].(string)
		switch action {
		case "CREATE_ROOM":
			return l.handleCreateRoomMessage(conn, userMessage)
		case "JOIN_ROOM":
			return l.handleJoinRoomMessage(conn, userMessage)
		case "LEAVE_ROOM":
			return l.handleLeaveRoomMessage(conn, userMessage)
		case "ROOM_STATE_CHANGE":
			return l.handleRoomStateChangeMessage(conn, userMessage)
		}
	} else {
		// Should never happen
	}

	return nil
}

func (l *Lobby) handleLoginMessage(conn *Connection, message map[string]interface{}) (err error) {
	// TODO: authentication
	delete(l.anonymousConn, conn)

	username := message["username"].(string)
	user := &User{
		id:       UserId(l.idGen.GetNextId()),
		state:    ACTIVE,
		conn:     conn,
		username: username,
		nickname: username,
	}
	l.userConnMap[conn] = user
	l.userStateMap[ACTIVE][user.id] = user

	l.BroadcastUserEnter(user)
	l.SyncLobbyState(user)

	log.Println("User login successful: userid = ", user.id, ", username = ", user.username)
	return nil
}

func (l *Lobby) handleCreateRoomMessage(conn *Connection, message map[string]interface{}) (err error) {
	if user, ok := l.userConnMap[conn]; ok {
		if _, ok := l.userInRoomMap[user.id]; ok {
			return nil
		}
		observerUsers := make(map[UserId]bool)
		observerUsers[user.id] = true
		room := &Room{
			id:            RoomId(l.idGen.GetNextId()),
			status:        RoomStatus(0),
			seatedUsers:   make([]UserId, 0),
			observerUsers: observerUsers,
		}
		l.rooms[room.id] = room
		l.userInRoomMap[user.id] = room

		l.BroadcastUserCreateRoom(user, room)
		l.SyncRoomState(user, room)
	}
	return nil
}

func (l *Lobby) handleJoinRoomMessage(conn *Connection, message map[string]interface{}) (err error) {
	if user, ok := l.userConnMap[conn]; ok {
		if _, ok := l.userInRoomMap[user.id]; ok {
			return nil
		}
		roomId := RoomId(message["room_id"].(float64))
		room := l.rooms[roomId]
		l.userInRoomMap[user.id] = room

		room.observerUsers[user.id] = true

		l.BroadcastUserJoinRoom(user, room)
		l.SyncRoomState(user, room)
	}
	return nil
}

func (l *Lobby) handleLeaveRoomMessage(conn *Connection, message map[string]interface{}) (err error) {
	if user, ok := l.userConnMap[conn]; ok {
		if _, ok := l.userInRoomMap[user.id]; !ok {
			return nil
		}
		roomId := message["room_id"].(RoomId)
		room := l.rooms[roomId]
		delete(room.observerUsers, user.id)
		l.userInRoomMap[user.id] = nil
		l.BroadcastUserLeaveRoom(user)
	}
	return nil
}

func (l *Lobby) handleRoomStateChangeMessage(conn *Connection, message map[string]interface{}) (err error) {
	if user, ok := l.userConnMap[conn]; ok {
		if _, ok := l.userInRoomMap[user.id]; !ok {
			return nil
		}
		roomId := message["room_id"].(RoomId)
		room := l.rooms[roomId]
		room.handleMessage(user, message)
	}
	return nil
}

func (l *Lobby) authenticate(username string, password string) bool {
	return true
}

func (l *Lobby) SyncLobbyState(user *User) {
	message := make(map[string]interface{})
	message["type"] = "SYNC_LOBBY_STATE"
	users := make([]map[string]interface{}, 0)
	rooms := make([]map[string]interface{}, 0)
	for _, m := range l.userStateMap {
		for _, u := range m {
			if u != user {
				user := make(map[string]interface{})
				user["id"] = u.id
				user["username"] = u.username
				user["nickname"] = u.nickname
				users = append(users, user)
			}
		}
	}
	for _, r := range l.rooms {
		room := r.serializeRoomAbstract()
		rooms = append(rooms, room)
	}
	message["users"] = users
	message["rooms"] = rooms

	str, err := json.Marshal(message)
	if err != nil {
		log.Println("SyncLobbyState error")
		return
	}
	user.conn.sendBuffer <- str
}

func (l *Lobby) SyncRoomState(user *User, room *Room) {
	message := make(map[string]interface{})
	message["type"] = "SYNC_ROOM_STATE"
	message["room"] = room.serializeRoom()

	str, err := json.Marshal(message)
	if err != nil {
		log.Println("SyncLobbyState error")
		return
	}
	user.conn.sendBuffer <- str
}

func (l *Lobby) BroadcastUserEnter(user *User) {
	message := make(map[string]interface{})
	message["type"] = "USER_ENTER_LOBBY"
	message["userid"] = user.id
	message["username"] = user.username
	message["nickname"] = user.nickname
	str, err := json.Marshal(message)
	if err != nil {
		log.Println("BroadcastUserEnterLobby error")
		return
	}
	for _, m := range l.userStateMap {
		for _, u := range m {
			if u != user {
				u.conn.sendBuffer <- str
			}
		}
	}
}

func (l *Lobby) BroadcastUserCreateRoom(user *User, room *Room) {
	message := make(map[string]interface{})
	message["type"] = "USER_CREATE_ROOM"
	message["user_id"] = user.id
	message["room_id"] = room.id
	str, err := json.Marshal(message)
	if err != nil {
		log.Println("BroadcastUserCreateRoom error")
		return
	}
	for _, m := range l.userStateMap {
		for _, u := range m {
			u.conn.sendBuffer <- str
		}
	}
}

func (l *Lobby) BroadcastUserJoinRoom(user *User, room *Room) {
	message := make(map[string]interface{})
	message["type"] = "USER_JOIN_ROOM"
	message["user_id"] = user.id
	message["room_id"] = room.id
	str, err := json.Marshal(message)
	if err != nil {
		log.Println("BroadcastUserJoinRoom error")
		return
	}
	for _, m := range l.userStateMap {
		for _, u := range m {
			u.conn.sendBuffer <- str
		}
	}
}

func (l *Lobby) BroadcastUserLeaveRoom(user *User) {
	message := make(map[string]interface{})
	message["type"] = "USER_LEAVE_ROOM"
	message["user_id"] = user.id
	str, err := json.Marshal(message)
	if err != nil {
		log.Println("BroadcastUserLeaveRoom error")
		return
	}
	for _, m := range l.userStateMap {
		for _, u := range m {
			u.conn.sendBuffer <- str
		}
	}
}

func (l *Lobby) BroadcastRoomStateChange() {
	// Room id, state
}
