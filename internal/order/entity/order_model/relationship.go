package order_model

type OrderItem struct {
	Id       int     `gorm:"column:id;"`
	OrderId  int     `json:"order_id" gorm:"column:order_id;"`
	Quantity int     `json:"quantity" gorm:"column:quanity;"`
	Note     *string `json:"note" gorm:"column:note;"`
}

func (OrderItem) TableName() string {
	return "order_item"
}

//
//// ==================================================== Restaurant ================================================================
//type Restaurant struct {
//	Id               int     `gorm:"column:id;"`
//	Name             string  `json:"name" gorm:"column:name;"`
//	Address          string  `json:"address" gorm:"column:address;"`
//	CityId           *int    `json:"cityId" gorm:"column:city_id;"`
//	Lat              float64 `json:"lat" gorm:"column:lat;"`
//	Lng              float64 `json:"lng" gorm:"column:lng;"`
//	ShippingFeePerKm float64 `json:"shippingFeePerKm" gorm:"column:shipping_fee_per_km;"`
//	Orders           []Order `json:"orders" gorm:"foreignKey:RestaurantId;references:Id;"`
//}
//
//func (Restaurant) TableName() string {
//	return "restaurant"
//}
//
//// ============================================= User ====================================================================
//type User struct {
//	Id        int     `gorm:"column:id;"`
//	FirstName string  `gorm:"column:first_name"`
//	LastName  string  `gorm:"column:last_name"`
//	Phone     string  `gorm:"column:phone"`
//	Orders    []Order `gorm:"foreignKey:UserId;references:Id;"`
//}
//
//func (User) TableName() string {
//	return "user"
//}
