package category_service

import (
	categorymodel "Food-Delivery/internal/category/model"
	"Food-Delivery/pkg/common"
	"context"
)

type CategoryRepository interface {
	Create(ctx context.Context, dto *categorymodel.CategoryCreateDto) error
	FindAllWithCondition(ctx context.Context, paging *common.Paging, query *categorymodel.QueryDTO, keys ...string) ([]categorymodel.Category, error)
	FindAllByIds(ctx context.Context, ids []int, keys ...string) ([]categorymodel.Category, error)
	FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*categorymodel.Category, error)
	UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *categorymodel.CategoryCreateDto) error
	DeleteDataWithCondition(ctx context.Context, condition map[string]any) error
}

type categoryService struct {
	cateRepo CategoryRepository
}

func NewCategoryService(cateRepo CategoryRepository) *categoryService {
	return &categoryService{cateRepo}
}

func (service *categoryService) Create(ctx context.Context, cate *categorymodel.CategoryCreateDto) error {
	//------perform business operation such as validate data
	if err := cate.Validate(); err != nil {
		return common.ErrBadRequest(err)
	}
	//------
	if err := service.cateRepo.Create(ctx, cate); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}

func (service *categoryService) FindAllByIds(ctx context.Context, ids []int) ([]categorymodel.Category, error) {
	//there will have business logic before getting data list with condition
	categories, err := service.cateRepo.FindAllByIds(ctx, ids)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return categories, nil
}

func (service *categoryService) FindAll(ctx context.Context, paging *common.Paging, filter *categorymodel.QueryDTO) ([]categorymodel.Category, error) {
	//there will have business logic before getting data list with condition
	categories, err := service.cateRepo.FindAllWithCondition(ctx, paging, filter)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return categories, nil
}

func (service *categoryService) FindOneById(ctx context.Context, id int) (*categorymodel.Category, error) {
	//there will have business logic before getting specific data with condition

	category, err := service.cateRepo.FindOneWithCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return category, nil
}

func (service *categoryService) Update(ctx context.Context, id int, dto *categorymodel.CategoryCreateDto) error {
	//validate the data first under this usecase layer
	if err := dto.Validate(); err != nil {
		return err
	}
	//check the eixstence of data in database
	if _, err := service.FindOneById(ctx, id); err != nil {
		return err
	}
	//_, err := service.cateRepo.FindDataWithCondition(ctx, map[string]any{"id": id})
	//
	//if err != nil {
	//	return common.ErrEntityNotFound(categorymodel.EntityName, err)
	//}

	if err := service.cateRepo.UpdateDataWithCondition(ctx, map[string]any{"id": id}, dto); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}

func (service *categoryService) Delete(ctx context.Context, id int) error {
	//check the existence of data in database
	if _, err := service.FindOneById(ctx, id); err != nil {
		return err
	}

	//if there is no returned error, we call the method DeleteDataByCondition of placeRepo interface
	if err := service.cateRepo.DeleteDataWithCondition(ctx, map[string]any{"id": id}); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}
