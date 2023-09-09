package entity

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PeonChatConfig struct {
	ID               uint           `gorm:"primarykey"`
	ChatId           int64          `gorm:"chat_id; index:chatid_un,unique"`
	Status           int            `gorm:"status; type:int2"`
	ChatName         string         `gorm:"chat_name"`
	Config           datatypes.JSON `gorm:"config"`
	JobConfig        datatypes.JSON `gorm:"job_conifg"`
	PermissionConifg datatypes.JSON `gorm:"permission_config"`
	AttachJson       datatypes.JSON `gorm:"attach_json"`
	CreatedTime      time.Time      `gorm:"created_time; type:timestamptz"`
}

func (PeonChatConfig) TableName() string {
	return "peon_chat_cfg"
}

func (m *PeonChatConfig) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedTime = time.Now().UTC()
	return nil
}
