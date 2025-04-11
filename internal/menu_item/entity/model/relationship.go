package menu_item_model

type Restaurant struct {
	Id               int        `json:"id" gorm:"column:id;"`
	Name             string     `json:"name" gorm:"column:name;"`
	Address          string     `json:"address" gorm:"column:address;"`
	CityId           *int       `json:"cityId" gorm:"column:city_id;"`
	Lat              float64    `json:"lat" gorm:"column:lat;"`
	Lng              float64    `json:"lng" gorm:"column:lng;"`
	ShippingFeePerKm float64    `json:"shippingFeePerKm" gorm:"column:shipping_fee_per_km;"`
	Status           string     `json:"status" gorm:"column:status;"`
	MenuItems        []MenuItem `json:"menu_items" gorm:"foreignKey:RestaurantId;references:Id;"`
}
