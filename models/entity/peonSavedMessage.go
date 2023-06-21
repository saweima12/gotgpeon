package entity

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PeonSavedMessage struct {
	gorm.Model
	ChatId      string         `gorm:"chat_id"`
	MessageId   string         `gorm:"message_id"`
	MessageJson datatypes.JSON `gorm:"message_json"`
	RecordDate  datatypes.Time `gorm:"record_date"`
}

func (PeonSavedMessage) TableName() string {
	return "peon_saved_message"
}
