package entity

import (
	"gorm.io/datatypes"
)

type PeonChatConfig struct {
	ID             uint           `gorm:"primarykey"`
	ChatId         string         `gorm:"chat_id, type:varchar(40)"`
	Status         string         `gorm:"status"`
	ChatName       string         `gorm:"chat_name"`
	ConfigJson     datatypes.JSON `gorm:"config_json"`
	PermissionJson datatypes.JSON `gorm:"permission_json"`
	AttachJson     datatypes.JSON `gorm:"attach_json"`
}

func (PeonChatConfig) TableName() string {
	return "peon_chat_config"
}
