package entity

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PeonBehaviorRecord struct {
	gorm.Model
	ChatId      string         `gorm:"chat_id"`
	UserId      string         `gorm:"user_id"`
	FullName    string         `gorm:"full_name"`
	MsgCount    int            `gorm:"msg_count"`
	MemberLevel int            `gorm:"member_level"`
	UpdateTime  datatypes.Time `gorm:"update_time"`
	CreatedTime datatypes.Time `gorm:"created_time"`
}

func (PeonBehaviorRecord) TableName() string {
	return "peon_behavior_record"
}
