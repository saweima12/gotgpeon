package models

import "time"

type MessageRecord struct {
	UserId      string    `json:"user_id"`
	Point       int       `json:"point"`
	MemberLevel int       `json:"member_level"`
	CreatedTime time.Time `json:"created_time"`
}
