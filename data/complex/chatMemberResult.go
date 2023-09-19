package complex

import "time"

type ChatMemberResult struct {
	ChatId     int64     `gorm:"column:chat_id"`
	MemberId   int64     `gorm:"column:member_id"`
	FullName   string    `gorm:"full_name"`
	MsgCount   int       `gorm:"column:msg_count; type:int4"`
	UpdateTime time.Time `gorm:"column:update_time; type:timestamptz"`
}
