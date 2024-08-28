package models

import "time"

type Room struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	UpdatedBy string `json:"updated_by"`
}

type RoomMember struct {
	ID       string    `json:"id"`
	RoomId   string    `json:"room_id"`
	UserId   string    `json:"user_id"`
	JoinedAt time.Time `json:"joined_at"`
	LeavedAt *time.Time `json:"leaved_at"`
}

type Message struct {
	ID string `json:"id"`
}

type RoomMessage struct {
	ID               string    `json:"id"`
	RoomID           string    `json:"room_id"`
	UserID           string    `json:"user_id"`
	Message          string    `json:"message"`
	Status           string    `json:"status"`
	MentionedUserIDs string    `json:"mentioned_user_ids"`
	ReplyMessagesID  string    `json:"reply_messages_id"`
	CreatedAt        time.Time `json:"created_at"`
}
