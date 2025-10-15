package model

type ItemOnOptional struct {
	ItemId     int `json:"item_id" gorm:"primaryKey" column:"item_id;not null"`
	OptionalId int `json:"optional_id" gorm:"primaryKey" column:"optional_id;not null"`
}

func (ItemOnOptional) TableName() string {
	return "item_on_optional"
}
