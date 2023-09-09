package entity

import "time"

type PeonMemberAllowlist struct {
	ID          uint      `gorm:"primarykey"`
	MemberId    int64     `gorm:"member_id"`
	Status      int8      `gorm:"status; type:int2"`
	CreatedTime time.Time `gorm:"created_time; type:timestamptz"`
}

func (PeonMemberAllowlist) TableName() string {
	return "peon_member_allowlist"
}
