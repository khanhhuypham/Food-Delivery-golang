package common

import (
	"gorm.io/gorm"
	"time"
)

type SQLModel struct {
	Id        int       `json:"id" gorm:"primaryKey;column:id"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
	//gorm.DeleteAt: use to perform soft deletion
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}
