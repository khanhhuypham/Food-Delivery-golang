package constant

type OrderStatus string

const (
	ORDER_PEDNING    OrderStatus = "Pending"
	ORDER_INPROGRESS OrderStatus = "In-Progress"
	ORDER_DELIVERED  OrderStatus = "Delivered"
	ORDER_COMPLETED  OrderStatus = "Completed"
	ORDER_CANCELLED  OrderStatus = "Cancelled"
)
