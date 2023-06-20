package entity

import "gorm.io/datatypes"

type PeonDeletedMessage struct {
	ChatId      string         `gorm:"chat_id"`
	ContentType string         `gorm:"content_type"`
	MessageJson datatypes.JSON `gorm:"message_json"`
	RecordDate  datatypes.Date `gorm:"record_date"`
}

func (PeonDeletedMessage) TableName() string {
	return "peon_deleted_message"
}
