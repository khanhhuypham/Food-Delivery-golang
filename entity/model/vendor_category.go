package model

import (
	vendor_category_dto "Food-Delivery/entity/dto/vendor_category"
	"Food-Delivery/pkg/common"
)

const VendorCategoryEntity = "vendor category"

type VendorCategory struct {
	common.SQLModel
	Image        *common.Image `json:"image" gorm:"column:image;"`
	Name         string        `json:"name" gorm:"column:name;not null;unique"`
	Description  *string       `json:"description" gorm:"column:description;"`
	Active       bool          `json:"active" gorm:"column:active;default:true"`
	RestaurantId int           `json:"restaurant_id" gorm:"column:restaurant_id;not null"`
	Items        []Item        `json:"items" gorm:"foreignKey:VendorCategoryId;references:Id"`
}

func (VendorCategory) TableName() string {
	return "vendor_category"
}

func (vendorCategory *VendorCategory) ToVendorCategoryDetailDTO() *vendor_category_dto.VendorCategoryDTO {
	dto := &vendor_category_dto.VendorCategoryDTO{
		ID:           vendorCategory.Id,
		Image:        vendorCategory.Image,
		Name:         vendorCategory.Name,
		Active:       vendorCategory.Active,
		Description:  vendorCategory.Description,
		RestaurantId: vendorCategory.RestaurantId,
		TotalItems:   len(vendorCategory.Items),
	}

	if len(vendorCategory.Items) > 0 {
		items := make([]vendor_category_dto.ItemDTO, 0, len(vendorCategory.Items)) // preallocate

		for _, item := range vendorCategory.Items {
			items = append(items, vendor_category_dto.ItemDTO{
				ID:               item.Id,
				Name:             item.Name,
				Image:            item.Image,
				Price:            item.Price,
				DeliveryTime:     item.DeliveryTime,
				CategoryId:       item.CategoryId,
				VendorCategoryId: item.VendorCategoryId,
				RestaurantId:     item.RestaurantId,
			})
		}
		dto.Items = items
	}

	return dto
}
