package models

import "time"

type Message struct {
	ID         int       `json:"id" db:"id"`
	SenderID   int       `json:"sender_id" db:"sender_id"`
	ReceiverID int       `json:"receiver_id" db:"receiver_id"`
	Text       string    `json:"text" db:"text"`
	Timestamp  time.Time `json:"timestamp" db:"timestamp"`
}
