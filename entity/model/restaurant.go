package model

import (
	"Food-Delivery/entity/constant"
	rating_dto "Food-Delivery/entity/dto/rating"
	restaurant_dto "Food-Delivery/entity/dto/restaurant"
	"Food-Delivery/pkg/common"
	"errors"
	"strings"
)

const RestaurantEntity = "restaurant"

type Restaurant struct {
	common.SQLModel
	OwnerId          int                       `json:"ownerId" gorm:"column:owner_id;"`
	Cover            *common.Image             `json:"cover" gorm:"column:cover;"`
	Logo             *common.Image             `json:"logo" gorm:"column:logo;"`
	Name             string                    `json:"name" gorm:"column:name;"`
	Email            *string                   `json:"email" gorm:"column:email;"`
	Phone            string                    `json:"phone" gorm:"column:phone;"`
	Address          string                    `json:"address" gorm:"column:address;"`
	CityId           *int                      `json:"cityId" gorm:"column:city_id;"`
	Lat              float64                   `json:"lat" gorm:"column:lat;"`
	Lng              float64                   `json:"lng" gorm:"column:lng;"`
	ShippingFeePerKm float64                   `json:"shippingFeePerKm" gorm:"column:shipping_fee_per_km;"`
	LikeCount        int                       `json:"like_count" gorm:"column:like_count; default:0"`
	Status           constant.RestaurantStatus `json:"status" gorm:"column:status;"`
	Description      *string                   `json:"description" gorm:"column:description;"`
	OperatingHours   *OperatingHours           `json:"operating_hours" gorm:"column:operating_hours;type:json"`
	Items            []Item                    `json:"items" gorm:"foreignKey:RestaurantId;references:Id;"`
	Orders           []Order                   `json:"orders" gorm:"foreignKey:RestaurantId;references:Id;"`
	VendorCategory   []VendorCategory          `json:"vendor_categories" gorm:"foreignKey:RestaurantId;references:Id;"`
	Rating           *Rating                   `json:"rating" gorm:"foreignKey:RestaurantId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type DailyHours struct {
	OpenTime  string `json:"open_time"`  // "08:00" in 24h format
	CloseTime string `json:"close_time"` // "22:00"
}
type OperatingHours struct {
	Monday    *DailyHours `json:"monday"`
	Tuesday   *DailyHours `json:"tuesday"`
	Wednesday *DailyHours `json:"wednesday"`
	Thursday  *DailyHours `json:"thursday"`
	Friday    *DailyHours `json:"friday"`
	Saturday  *DailyHours `json:"saturday"`
	Sunday    *DailyHours `json:"sunday"`
}

func (Restaurant) TableName() string {
	return "restaurant"
}

func (r *Restaurant) Validate() error {
	r.Name = strings.TrimSpace(r.Name)

	r.Address = strings.TrimSpace(r.Address)

	if r.Name == "" {
		return common.ErrBadRequest(errors.New("restaurant name is required"))
	}

	if r.Address == "" {
		return common.ErrBadRequest(errors.New("restaurant address is required"))
	}

	if !r.Status.IsValid() {
		return common.ErrBadRequest(errors.New("invalid restaurant status"))
	}

	return nil
}

func (restaurant *Restaurant) ToRestaurantDTO() *restaurant_dto.RestaurantDTO {
	dto := &restaurant_dto.RestaurantDTO{
		ID:          restaurant.Id,
		Name:        restaurant.Name,
		Status:      restaurant.Status,
		Description: restaurant.Description,
	}

	if restaurant.Rating != nil {
		dto.Rating = &rating_dto.RatingDTO{
			Like:  restaurant.Rating.Like,
			Score: restaurant.Rating.Score,
		}
	}

	return dto
}
