package models

const (
	OK      = "ok"
	NG      = "ng"
	UNKNOWN = "unknown"
)

const (
	BLOCK  = -1
	NONE   = 0
	LIMIT  = 1
	JUNIOR = 2
	SENIOR = 3
)

var MemberLevelMap map[string]int = map[string]int{
	"block": BLOCK,
	"none":  NONE,
	"limit": LIMIT,
	"jr":    JUNIOR,
	"sr":    SENIOR,
}
