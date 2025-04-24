package category_service

import (
	category_dto "Food-Delivery/entity/dto/category"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
)

type CategoryRepository interface {
	Create(ctx context.Context, dto *category_dto.CreateDto) error
	FindAllWithCondition(ctx context.Context, paging *common.Paging, query *category_dto.QueryDTO, keys ...string) ([]model.Category, error)
	FindAllByIds(ctx context.Context, ids []int, keys ...string) ([]model.Category, error)
	FindOneWithCondition(ctx context.Context, condition map[string]any, keys ...string) (*model.Category, error)
	UpdateDataWithCondition(ctx context.Context, condition map[string]any, dto *category_dto.CreateDto) error
	DeleteDataWithCondition(ctx context.Context, condition map[string]any) error
}

//type MediaService interface {
//	Delete(ctx context.Context, id int)
//}

type categoryService struct {
	cateRepo CategoryRepository
	//mediaService MediaService
}

func NewCategoryService(cateRepo CategoryRepository) *categoryService {
	return &categoryService{cateRepo}
}

func (service *categoryService) Create(ctx context.Context, cate *category_dto.CreateDto) error {
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

func (service *categoryService) FindAllByIds(ctx context.Context, ids []int) ([]model.Category, error) {
	//there will have business logic before getting data list with condition
	categories, err := service.cateRepo.FindAllByIds(ctx, ids)

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return categories, nil
}

func (service *categoryService) FindAll(ctx context.Context, paging *common.Paging, filter *category_dto.QueryDTO) ([]model.Category, error) {
	//there will have business logic before getting data list with condition
	categories, err := service.cateRepo.FindAllWithCondition(ctx, paging, filter, "Items")

	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}

	return categories, nil
}

func (service *categoryService) FindOneById(ctx context.Context, id int) (*model.Category, error) {
	//there will have business logic before getting specific data with condition

	category, err := service.cateRepo.FindOneWithCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return nil, common.ErrInternal(err).WithDebug(err.Error())
	}
	return category, nil
}

func (service *categoryService) Update(ctx context.Context, id int, dto *category_dto.CreateDto) error {
	//validate the data first under this usecase layer
	if err := dto.Validate(); err != nil {
		return err
	}
	//check the eixstence of data in database
	if _, err := service.FindOneById(ctx, id); err != nil {
		return err
	}

	//if dto.Image != nil {
	//	service.mediaService.Delete(ctx, dto.Image.Id)
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
