package entity

import (
	"time"

	"gorm.io/gorm"
)

type PeonChatMemberRecord struct {
	ID          uint      `gorm:"primarykey"`
	ChatId      int64     `gorm:"column:chat_id; index:chatid_memberid_index,priority:1,unique"`
	MemberId    int64     `gorm:"column:member_id; index:chatid_memberid_index,priority:2,unique"`
	MsgCount    int       `gorm:"column:msg_count; type:int4"`
	MemberLevel int       `gorm:"column:member_level; type:int2"`
	UpdateTime  time.Time `gorm:"column:update_time; type:timestamptz"`
	CreatedTime time.Time `gorm:"column:created_time; type:timestamptz"`
}

func (PeonChatMemberRecord) TableName() string {
	return "peon_chat_member_record"
}

func (m *PeonChatMemberRecord) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedTime = time.Now().UTC()
	m.UpdateTime = time.Now().UTC()
	return nil
}

func (m *PeonChatMemberRecord) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdateTime = time.Now().UTC()
	return nil
}
