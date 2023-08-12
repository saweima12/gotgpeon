package entity

import (
	"time"

	"gorm.io/datatypes"
)

type PeonSavedMessage struct {
	ID          uint           `gorm:"primarykey"`
	ChatId      int64          `gorm:"chat_id"`
	MemberId    int64          `gorm:"member_id"`
	MessageJson datatypes.JSON `gorm:"message_json"`
	RecordDate  time.Time      `gorm:"record_date; type:timestamptz"`
}

func (PeonSavedMessage) TableName() string {
	return "peon_saved_msg"
}
