package entity

import (
	"time"

	"gorm.io/gorm"
)

type PeonMemberRecord struct {
	ID          uint      `gorm:"primarykey"`
	MemberId    int64     `gorm:"member_id; index:member_idx,unique"`
	FullName    string    `gorm:"full_name"`
	CreatedTime time.Time `gorm:"created_time; type:timestamptz"`
}

func (PeonMemberRecord) TableName() string {
	return "peon_member_record"
}

func (m *PeonMemberRecord) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedTime = time.Now().UTC()
	return nil
}
