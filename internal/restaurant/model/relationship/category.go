package relationship

type Category struct {
	Id     int    `json:"id" gorm:"column:id;"`
	Name   string `json:"name" gorm:"column:name;"`
	Status string `json:"status" gorm:"column:status;"`
}

func (Category) TableName() string {
	return "category"
}
