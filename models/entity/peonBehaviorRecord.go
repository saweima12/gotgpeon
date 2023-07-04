package entity

import (
	"time"

	"gorm.io/gorm"
)

type PeonBehaviorRecord struct {
	ID          uint      `gorm:"primarykey"`
	ChatId      string    `gorm:"column:chat_id; type:varchar(40); unique"`
	UserId      string    `gorm:"column:user_id; type:varchar(40); unique"`
	FullName    string    `gorm:"column:full_name; type:text"`
	MsgCount    int       `gorm:"column:msg_count; type:int4"`
	MemberLevel int       `gorm:"column:member_level; type:int2"`
	UpdateTime  time.Time `gorm:"column:update_time"`
	CreatedTime time.Time `gorm:"column:created_time"`
}

func (PeonBehaviorRecord) TableName() string {
	return "peon_behavior_record"
}

func (m *PeonBehaviorRecord) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedTime = time.Now()
	m.UpdateTime = time.Now()
	return nil
}

func (m *PeonBehaviorRecord) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdateTime = time.Now()
	return nil
}
