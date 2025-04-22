package driver_service

type DriverRepository interface {
}

type driverService struct {
	driverRepo DriverRepository
}

func NewOrderService(driverRepo DriverRepository) *driverService {
	return &driverService{driverRepo}
}
