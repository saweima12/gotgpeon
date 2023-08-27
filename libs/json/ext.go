package json

import jsoniter "github.com/json-iterator/go"

var (
	cjson               = jsoniter.ConfigCompatibleWithStandardLibrary
	Unmarshal           = cjson.Unmarshal
	UnmarshalFromString = cjson.UnmarshalFromString
	Marshal             = cjson.Marshal
	MarshalToString     = cjson.MarshalToString
)
