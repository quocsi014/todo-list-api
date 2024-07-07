package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
)

type ItemStatus int

const (
	Todo ItemStatus = iota
	Doing
	Done
)

var StatusStrValue [3]string = [3]string{"To do", "Doing", "Done"}

func (item *ItemStatus) String() string {
	return StatusStrValue[*item]
}

func ParseStrToStatus(str string, item *ItemStatus) error {
	for i, v := range StatusStrValue {
		if v == str {
			*item = ItemStatus(i)
			return nil
		}
	}
	return errors.New("invalid string value")
}

func (item *ItemStatus) Scan(value interface{}) error {
	byteValue, ok := value.([]byte)

	if !ok {
		return errors.New("fail to scan data from sql")
	}

	strValue := string(byteValue)

	return ParseStrToStatus(strValue, item)
}

func (item *ItemStatus) Value() (driver.Value, error) {
	if *item > 2 {
		return nil, errors.New("ItemStatus is not invalid")
	}
	return item.String(), nil
}

func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(StatusStrValue[*item])
}

func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	return ParseStrToStatus(str, item)
}
