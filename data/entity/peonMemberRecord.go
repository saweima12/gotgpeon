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
	UpdatedTime time.Time `gorm:"updated_time; type:timestamptz"`
}

func (PeonMemberRecord) TableName() string {
	return "peon_member_record"
}

func (m *PeonMemberRecord) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedTime = time.Now().UTC()
	m.UpdatedTime = time.Now().UTC()
	return nil
}

func (m *PeonMemberRecord) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedTime = time.Now().UTC()
	return nil
}
