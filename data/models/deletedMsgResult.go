package models

import "github.com/goccy/go-json"

type DeletedMessage struct {
	ContentType string
	Content     json.RawMessage
	RecordTime  int64
}
