package constant

type CategoryStatus string

const (
	CATEGORY_ACTIVE   = "active"
	CATEGORY_INACTIVE = "inactive"
	CATEGORY_DELETED  = "deleted"
)

// Method to check if the status is valid
func (s CategoryStatus) IsValid() bool {
	switch s {
	case CATEGORY_ACTIVE, CATEGORY_INACTIVE, CATEGORY_DELETED:
		return true
	}
	return false
}
