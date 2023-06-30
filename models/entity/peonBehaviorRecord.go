package entity

import (
	"time"
)

type PeonBehaviorRecord struct {
	ChatId      string    `gorm:"chat_id"`
	UserId      string    `gorm:"user_id"`
	FullName    string    `gorm:"full_name"`
	MsgCount    int       `gorm:"msg_count"`
	MemberLevel int       `gorm:"member_level"`
	UpdateTime  time.Time `gorm:"update_time"`
	CreatedTime time.Time `gorm:"created_time"`
}

func (PeonBehaviorRecord) TableName() string {
	return "peon_behavior_record"
}
