package entity

import (
	"gorm.io/datatypes"
)

type PeonUserWhitelist struct {
	ID         uint           `gorm:"primarykey"`
	UserId     string         `gorm:"user_id"`
	Status     string         `gorm:"status"`
	CreateDate datatypes.Time `gorm:"create_date"`
}

func (PeonUserWhitelist) TableName() string {
	return "peon_user_whitelist"
}
