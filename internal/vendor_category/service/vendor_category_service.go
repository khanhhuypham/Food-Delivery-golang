package vendor_category_service

import (
	vendor_category_dto "Food-Delivery/entity/dto/vendor_category"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"errors"
	"gorm.io/gorm"
)

type VendorCategoryRepository interface {
	FindAllByRestaurantId(ctx context.Context, restaurantId int, keys ...string) ([]model.VendorCategory, error)
	FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.VendorCategory, error)
	Create(ctx context.Context, dto *vendor_category_dto.CreateDTO) (*model.VendorCategory, error)
	UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *vendor_category_dto.UpdateDTO) (*model.VendorCategory, error)
	DeleteDataWithCondition(ctx context.Context, condition map[string]any) error
}

type vendorCategoryService struct {
	vendorCategoryRepo VendorCategoryRepository
}

func NewVendorCategoryService(vendorCategoryRepo VendorCategoryRepository) *vendorCategoryService {
	return &vendorCategoryService{vendorCategoryRepo}
}

func (service *vendorCategoryService) FindAll(ctx context.Context, restaurantId int) ([]model.VendorCategory, error) {
	//there will have business logic before getting data list with condition
	items, err := service.vendorCategoryRepo.FindAllByRestaurantId(ctx, restaurantId, "Items")

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return items, nil
}

func (service *vendorCategoryService) FindOneById(ctx context.Context, id int) (*model.VendorCategory, error) {
	//there will have business logic before getting specific data with condition

	data, err := service.vendorCategoryRepo.FindOneWithCondition(ctx, map[string]any{"id": id}, "Items")
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound(model.ItemEntity, err).WithDebug(err.Error())
		}
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return data, nil
}

func (service *vendorCategoryService) Create(ctx context.Context, dto *vendor_category_dto.CreateDTO) (*model.VendorCategory, error) {
	//------perform business operation such as validate data
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	//------
	data, err := service.vendorCategoryRepo.Create(ctx, dto)
	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return data, nil
}

func (service *vendorCategoryService) Update(ctx context.Context, id int, dto *vendor_category_dto.UpdateDTO) (*model.VendorCategory, error) {

	if err := dto.Validate(); err != nil {
		return nil, err
	}

	if _, err := service.FindOneById(ctx, id); err != nil {
		return nil, err
	}

	updatedData, err := service.vendorCategoryRepo.UpdateDataWithCondition(ctx, map[string]any{"id": id}, dto)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return updatedData, nil
}

func (service *vendorCategoryService) Delete(ctx context.Context, id int) error {

	if _, err := service.FindOneById(ctx, id); err != nil {
		return err
	}

	//if there is no returned error, we call the method DeleteDataByCondition of placeRepo interface
	if err := service.vendorCategoryRepo.DeleteDataWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}
