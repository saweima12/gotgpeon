package utils

import (
	"gotgpeon/logger"
	"gotgpeon/utils/jsonutil"
)

func AnyToStruct(data interface{}, dest any) error {

	bytes, err := jsonutil.Marshal(data)
	if err != nil {
		logger.Errorf("InterfaceToStruct Marshal err :%s", err.Error())
		return err
	}

	err = jsonutil.Unmarshal(bytes, dest)
	if err != nil {
		logger.Errorf("InterfaceToStruct UnMarshal err: %s", err.Error())
		return err
	}

	return nil
}
