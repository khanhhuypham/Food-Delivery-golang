package constant

type RestaurantStatus int

const (
	RESTAURANT_STATUS_OPEN                 RestaurantStatus = 1 //- The store is currently operating and accepting orders.
	RESTAURANT_STATUS_CLOSED               RestaurantStatus = 2 // - The store is not operating (e.g., outside business hours)
	RESTAURANT_STATUS_TEMPORARILY_CLOSED   RestaurantStatus = 3 //– Closed due to temporary reasons (e.g., holiday, maintenance).
	RESTAURANT_STATUS_LIMITED_AVAILABILITY RestaurantStatus = 4 // -Temporarily not accepting orders due to high load
	RESTAURANT_STATUS_SUSPENDED            RestaurantStatus = 5 //→ Admin-disabled due to issues (like payment or policy)
)

func (status RestaurantStatus) IsValid() bool {
	switch status {
	case
		RESTAURANT_STATUS_OPEN,
		RESTAURANT_STATUS_CLOSED,
		RESTAURANT_STATUS_TEMPORARILY_CLOSED,
		RESTAURANT_STATUS_LIMITED_AVAILABILITY,
		RESTAURANT_STATUS_SUSPENDED:
		return true
	}
	return false
}
