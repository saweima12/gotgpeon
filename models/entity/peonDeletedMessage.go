package entity

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PeonDeletedMessage struct {
	ID          uint           `gorm:"primarykey"`
	ChatId      int64          `gorm:"chat_id; index:chatid"`
	ContentType string         `gorm:"content_type"`
	MessageJson datatypes.JSON `gorm:"message_json"`
	RecordTime  time.Time      `gorm:"record_time; type:timestamptz"`
}

func (PeonDeletedMessage) TableName() string {
	return "peon_deleted_msg"
}

func (m *PeonDeletedMessage) BeforeCreate(tx *gorm.DB) (err error) {
	m.RecordTime = time.Now().UTC()
	return nil
}
