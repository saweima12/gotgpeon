package models

import "time"

type MessageRecord struct {
	MemberId    int64     `json:"member_id"`
	FullName    string    `json:"user_name"`
	Point       int       `json:"point"`
	MemberLevel int       `json:"member_level"`
	CreatedTime time.Time `json:"created_time"`
}

func NewMessageRecord(userId int64, fullName string) *MessageRecord {
	return &MessageRecord{
		MemberId:    userId,
		FullName:    fullName,
		Point:       0,
		MemberLevel: 0,
		CreatedTime: time.Now(),
	}
}
