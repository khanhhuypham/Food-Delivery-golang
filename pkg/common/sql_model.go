package common

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type SQLModel struct {
	Id        int      `json:"id" gorm:"primaryKey;column:id"`
	CreatedAt JSONTime `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt JSONTime `json:"updatedAt" gorm:"column:updated_at"`
	//gorm.DeleteAt: use to perform soft deletion
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

//func (t *JSONTime) UnmarshalJSON(b []byte) error {
//	s := strings.Trim(string(b), `"`)
//	parsed, err := time.Parse("2006-01-02 15:04:05", s)
//	if err != nil {
//		return err
//	}
//	*t = JSONTime(parsed)
//	return nil
//}

// Value implements the driver.Valuer interface for saving to DB
func (t JSONTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

// Scan implements the sql.Scanner interface for reading from DB
func (t *JSONTime) Scan(value interface{}) error {
	if value == nil {
		*t = JSONTime(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*t = JSONTime(v)
		return nil
	case []byte:
		parsed, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			return err
		}
		*t = JSONTime(parsed)
		return nil
	case string:
		parsed, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return err
		}
		*t = JSONTime(parsed)
		return nil
	default:
		return fmt.Errorf("unsupported Scan type for JSONTime: %T", value)
	}
}
