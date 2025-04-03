package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id     int    `json:"id" gorm:"column:id"`
	Url    string `json:"url" gorm:"column:url"`
	Width  int    `json:"width" gorm:"column:width"`
	Height int    `json:"height" gorm:"column:height"`
}

func (Image) TableName() string {
	return "image"
}

func (i *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var image Image
	if err := json.Unmarshal(bytes, &image); err != nil {
		return err
	}
	*i = image
	return nil
}

func (i *Image) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}
	return json.Marshal(i)
}
