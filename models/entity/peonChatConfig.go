package entity

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PeonChatConfig struct {
	gorm.Model
	ChatId         string         `gorm:"chat_id"`
	Status         string         `gorm:"status"`
	ChatName       string         `gorm:"chat_name"`
	ConfigJson     datatypes.JSON `gorm:"config_json"`
	PermissionJson datatypes.JSON `gorm:"permission_json"`
	AttachJson     datatypes.JSON `gorm:"attach_json"`
}

func (PeonChatConfig) TableName() string {
	return "peon_chat_config"
}
