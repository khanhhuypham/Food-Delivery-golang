package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

const MediaEntity = "media"

type Media struct {
	Id        int    `json:"id" gorm:"column:id"`
	Url       string `json:"url" gorm:"-"`
	Filename  string `json:"-" gorm:"column:filename;"`
	Folder    string `json:"-" gorm:"column:folder;"`
	CloudName string `json:"-" gorm:"column:cloud_name;"`
	Size      int64  `json:"size" gorm:"column:size;"`
	Height    *int   `json:"height" gorm:"column:height;"`
	Width     *int   `json:"width" gorm:"column:width;"`
	Ext       string `json:"-" gorm:"column:ext;"`
	Status    string `json:"-" gorm:"column:status;"`
}

func (Media) TableName() string {
	return "media"
}

func (m *Media) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var media Media
	if err := json.Unmarshal(bytes, &media); err != nil {
		return err
	}
	*m = media
	return nil
}

func (m *Media) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}
