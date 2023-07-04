package models

import "time"

type MessageRecord struct {
	UserId      string    `json:"user_id"`
	FullName    string    `json:"user_name"`
	Point       int       `json:"point"`
	MemberLevel int       `json:"member_level"`
	CreatedTime time.Time `json:"created_time"`
}

func NewMessageRecord(userId string, fullName string) *MessageRecord {
	return &MessageRecord{
		UserId:      userId,
		FullName:    fullName,
		Point:       0,
		MemberLevel: 0,
		CreatedTime: time.Now(),
	}
}
