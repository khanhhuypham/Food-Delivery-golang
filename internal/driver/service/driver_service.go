package driver_service

import (
	driver_dto "Food-Delivery/entity/dto/driver"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"errors"
	"gorm.io/gorm"
)

type DriverRepository interface {
	Create(ctx context.Context, dto *driver_dto.CreateDTO) (*model.Driver, error)
	FindAllWithCondition(
		ctx context.Context,
		paging *common.Paging,
		query *driver_dto.QueryDTO,
		keys ...string) ([]model.Driver, error)

	FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.Driver, error)
	DeleteDataWithCondition(ctx context.Context, condition map[string]any) error
	UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *driver_dto.CreateDTO) (*model.Driver, error)
}
type driverService struct {
	driverRepo DriverRepository
}

func NewDriverService(driverRepo DriverRepository) *driverService {
	return &driverService{driverRepo}
}

func (service *driverService) Create(ctx context.Context, cate *driver_dto.CreateDTO) (*model.Driver, error) {
	//------perform business operation such as validate data
	if err := cate.Validate(); err != nil {
		return nil, common.ErrBadRequest(err)
	}
	//------
	driver, err := service.driverRepo.Create(ctx, cate)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return driver, nil
}

//
//func (service *driverService) FindAllByIds(ctx context.Context, ids []int) ([]model.Category, error) {
//	//there will have business logic before getting data list with condition
//	//categories, err := service.driverRepo.FindAllByIds(ctx, ids)
//
//	if err != nil {
//		return nil, common.ErrInternal(err).WithDebug(err.Error())
//	}
//
//	return categories, nil
//}

func (service *driverService) FindAll(ctx context.Context, paging *common.Paging, filter *driver_dto.QueryDTO) ([]model.Driver, error) {
	//there will have business logic before getting data list with condition
	list, err := service.driverRepo.FindAllWithCondition(ctx, paging, filter)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return list, nil
}

func (service *driverService) FindOneById(ctx context.Context, id int) (*model.Driver, error) {
	//there will have business logic before getting specific data with condition
	driver, err := service.driverRepo.FindOneWithCondition(ctx, map[string]any{"id": id})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrNotFound(err).WithDebug(err.Error())
		}
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return driver, nil
}

func (service *driverService) Update(ctx context.Context, id int, dto *driver_dto.CreateDTO) (*model.Driver, error) {
	//validate the data first under this usecase layer
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	//check the eixstence of data in database
	if _, err := service.FindOneById(ctx, id); err != nil {
		return nil, err
	}

	//if dto.Image != nil {
	//	service.mediaService.Delete(ctx, dto.Image.Id)
	//}

	driver, err := service.driverRepo.UpdateDataWithCondition(ctx, map[string]any{"id": id}, dto)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return driver, nil
}

func (service *driverService) Delete(ctx context.Context, id int) error {
	//check the existence of data in database
	if _, err := service.FindOneById(ctx, id); err != nil {
		return err
	}

	//if there is no returned error, we call the method DeleteDataByCondition of placeRepo interface
	if err := service.driverRepo.DeleteDataWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}
