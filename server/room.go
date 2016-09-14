package main

import (
	"encoding/json"
	"log"

	"hubble/gamemod"
)

type RoomStatus int8

const (
	WAITING RoomStatus = iota
	PLAYING
)

type Room struct {
	lobby *Lobby

	id      RoomId
	status  RoomStatus
	gameMod GameMod

	seatedUsers   []UserId
	observerUsers map[UserId]bool
}

func (r *Room) serializeRoom() (result map[string]interface{}) {
	result = make(map[string]interface{})
	result["id"] = r.id
	result["status"] = r.status
	result["seated_users"] = make(map[int64]interface{})
	for i, u := range r.seatedUsers {
		result["seated_users"].(map[int64]interface{})[int64(i)] = int64(u)
	}
	result["observer_users"] = make([]int64, 0)
	for u, _ := range r.observerUsers {
		result["observer_users"] = append(result["observer_users"].([]int64), int64(u))
	}
	return
}

func (r *Room) serializeRoomAbstract() (result map[string]interface{}) {
	result = make(map[string]interface{})
	result["id"] = r.id
	result["status"] = r.status

	allAttendees := make([]int64, 0)
	for _, u := range r.seatedUsers {
		allAttendees = append(allAttendees, int64(u))
	}
	for u, _ := range r.observerUsers {
		allAttendees = append(allAttendees, int64(u))
	}
	result["attendees"] = allAttendees
	return
}

func (r *Room) handleMessage(user *User, message map[string]interface{}) (err error) {
	action := message["room_action"].(string)
	switch action {
	case "TO_SEAT":
		return r.handleToSeatMessage(user, message)
	case "LEAVE_SEAT":
		return r.handleLeaveSeatMessage(user, message)
	case "START_PLAYING":
		return r.handleStartPlayingMessage(user, message)
	case "GAME":
		return r.handleGameMessage(user, message)
	}
	return nil
}

func (r *Room) handleToSeatMessage(user *User, message map[string]interface{}) (err error) {
	if r.status != WAITING {
		return nil
	}
	seat_id := int64(message["seat_id"].(float64))
	if r.seatedUsers[seat_id] == user.id || r.seatedUsers[seat_id] != 0 {
		return nil
	}

	for i, u := range r.seatedUsers {
		if u == user.id {
			r.seatedUsers[i] = UserId(0)
		}
	}

	delete(r.observerUsers, user.id)
	r.seatedUsers[seat_id] = user.id

	m := make(map[string]interface{})
	m["type"] = "TO_SEAT"
	m["user_id"] = user.id
	m["seat_id"] = seat_id
	r.BroadcastRoomStateChange(m)
	return nil
}

func (r *Room) handleLeaveSeatMessage(user *User, message map[string]interface{}) (err error) {
	if r.status != WAITING {
		return nil
	}

	seat_id := int64(-1)
	for i, u := range r.seatedUsers {
		if u == user.id {
			r.seatedUsers[i] = UserId(0)
			seat_id = int64(i)
			break
		}
	}
	if seat_id < 0 {
		return nil
	}
	m := make(map[string]interface{})
	m["type"] = "LEAVE_SEAT"
	m["user_id"] = user.id
	r.BroadcastRoomStateChange(m)
	return nil
}

func (r *Room) handleStartPlayingMessage(user *User, message map[string]interface{}) (err error) {
	if r.status != WAITING {
		return nil
	}
	r.status = PLAYING
	m := make(map[string]interface{})
	m["type"] = "START_PLAYING"
	r.BroadcastRoomStateChange(m)
	return nil
}

func (r *Room) handleGameMessage(user *User, message map[string]interface{}) (err error) {
	return nil
}

func (r *Room) BroadcastRoomStateChange(message map[string]interface{}) {
	str, err := json.Marshal(message)
	if err != nil {
		log.Println("BroadcastRoomStateChange error")
		return
	}
	for _, u := range r.seatedUsers {
		if user, ok := r.lobby.userIdMap[u]; ok {
			user.conn.sendBuffer <- str
		}
	}
	for u, _ := range r.observerUsers {
		if user, ok := r.lobby.userIdMap[u]; ok {
			user.conn.sendBuffer <- str
		}
	}
}
