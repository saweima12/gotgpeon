package utils

import (
	"gotgpeon/libs/json"
	"gotgpeon/logger"
)

func AnyToStruct(data interface{}, dest any) error {

	bytes, err := json.Marshal(data)
	if err != nil {
		logger.Errorf("InterfaceToStruct Marshal err :%s", err.Error())
		return err
	}

	err = json.Unmarshal(bytes, &dest)
	if err != nil {
		logger.Errorf("InterfaceToStruct UnMarshal err: %s", err.Error())
		return err
	}

	return nil
}
