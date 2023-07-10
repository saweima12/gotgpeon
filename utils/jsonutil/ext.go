package jsonutil

import jsoniter "github.com/json-iterator/go"

var (
	json                = jsoniter.ConfigCompatibleWithStandardLibrary
	Unmarshal           = json.Unmarshal
	UnmarshalFromString = json.UnmarshalFromString
	Marshal             = json.Marshal
	MarshalToString     = json.MarshalToString
)
