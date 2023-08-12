package entity

type PeonMemberRecord struct {
	ID       uint  `gorm:"primarykey"`
	MemberId int64 `gorm:"member_id"`
}

func (PeonMemberRecord) TableName() string {
	return "peon_member_record"
}
