package entity

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PeonUserWhitelist struct {
	gorm.Model
	UserId     string         `gorm:"user_id"`
	Status     string         `gorm:"status"`
	CreateDate datatypes.Time `gorm:"create_date"`
}

func (PeonUserWhitelist) TableName() string {
	return "peon_user_whitelist"
}
