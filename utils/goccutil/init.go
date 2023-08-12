package goccutil

import (
	"gotgpeon/logger"

	"github.com/liuzl/gocc"
)

var s2t *gocc.OpenCC

func InitOpenCC() error {
	var err error
	s2t, err = gocc.New("s2t")

	if err != nil {
		logger.Error("OpenCC intialize error: " + err.Error())
		return err
	}

	return nil
}

func S2T(data string) string {
	newStr, err := s2t.Convert(data)
	if err != nil {
		logger.Errorf("S2T error %s", err.Error())
	}
	return newStr
}
