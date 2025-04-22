package common

const (
	DBTypeUser          = 1
	DBTypePlace         = 2
	DBTypeLocation      = 3
	DBTypeAmenity       = 4
	DBTypePlace_Amenity = 5
	DBTypeBooking       = 6
	DBTypeReview        = 7
	DBTypeCity          = 8

	KeyRequester = "requester"
)

type Requester interface {
	//GetUserId() int
	GetUserEmail() string
	GetFullName() string
	GetUserRole() string
}

const (
	EventUserLikeRestaurant    = "UserLikedRestaurant"
	EventUserDislikeRestaurant = "UserUnlikedRestaurant"
)
