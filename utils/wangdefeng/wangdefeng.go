package wangdefeng

import (
	"encoding/json"
	"errors"
	"fmt"
)

func Convert2Target(data, target interface{}) error {
	encodeData, err := json.Marshal(data)
	if err != nil {
		return errors.New(fmt.Sprintf("marshal error: %s", err.Error()))
	}
	err = json.Unmarshal(encodeData, target)
	if err != nil {
		return errors.New(fmt.Sprintf("unmarshal error: %s", err.Error()))
	}
	return err
}

func Json2Target(data string, target interface{}) error {
	bytes := []byte(data)
	err := json.Unmarshal(bytes, target)
	if err != nil {
		return errors.New(fmt.Sprintf("unmarshal error: %s", err.Error()))
	}
	return err
}
