package entity

import (
	"gorm.io/datatypes"
)

type PeonSavedMessage struct {
	ChatId      string         `gorm:"chat_id"`
	MessageId   string         `gorm:"message_id"`
	MessageJson datatypes.JSON `gorm:"message_json"`
	RecordDate  datatypes.Time `gorm:"record_date"`
}

func (PeonSavedMessage) TableName() string {
	return "peon_saved_message"
}
