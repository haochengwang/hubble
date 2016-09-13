package main

import (
	"encoding/json"
	"log"
)

type RoomStatus int8

const (
	WAITING RoomStatus = iota
	PLAYING
)

type Room struct {
	lobby *Lobby

	id            RoomId
	status        RoomStatus
	seatedUsers   []UserId
	observerUsers map[UserId]bool
}

func (r *Room) serializeRoom() (result map[string]interface{}) {
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
	return nil
}

func (r *Room) handleLeaveSeatMessage(user *User, message map[string]interface{}) (err error) {
	return nil
}

func (r *Room) handleStartPlayingMessage(user *User, message map[string]interface{}) (err error) {
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
