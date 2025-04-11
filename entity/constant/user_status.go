package constant

type UserType string
type UserRole string
type UserStatus string

const (
	TypeEmailPassword UserType = "email_password"
	TypeFacebook      UserType = "facebook"
	TypeGmail         UserType = "gmail"

	ROLE_USER     UserRole = "user"
	ROLE_ADMIN    UserRole = "admin"
	ROLE_SHIPPPER UserRole = "shipper"

	USER_PENDING  UserStatus = "pending"
	USER_ACTIVE   UserStatus = "active"
	USER_INACTIVE UserStatus = "inactive"
	USER_BANNED   UserStatus = "banned"
	USER_DELETED  UserStatus = "deleted"
)
