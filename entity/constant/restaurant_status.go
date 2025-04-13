package constant

type RestaurantStatus string

const (
	RESTAURANT_ACTIVE   RestaurantStatus = "active"
	RESTAURANT_INACTIVE RestaurantStatus = "inactive"
	RESTAURANT_PENDING  RestaurantStatus = "pending"
	RESTAURANT_DELETED  RestaurantStatus = "deleted"
)

func (s RestaurantStatus) IsValid() bool {
	switch s {
	case RESTAURANT_ACTIVE, RESTAURANT_INACTIVE, RESTAURANT_PENDING, RESTAURANT_DELETED:
		return true
	}
	return false
}
