package rating_service

import (
	rating_dto "Food-Delivery/entity/dto/rating"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
)

type RatingRepository interface {
	Create(ctx context.Context, dto *rating_dto.CreateDTO) error
	Update(ctx context.Context, id int, dto *rating_dto.CreateDTO) error
	DeleteLike(ctx context.Context, condition map[string]any) error
	FindOneWithCondition(ctx context.Context, condition map[string]any) (*model.Rating, error)
}

type EventPublisher interface {
	Publish(ctx context.Context, topic string, event *common.AppEvent) error
}

type ratingService struct {
	repo           RatingRepository
	eventPublisher EventPublisher
}

func NewRatingService(repo RatingRepository, event EventPublisher) *ratingService {
	return &ratingService{
		repo:           repo,
		eventPublisher: event,
	}
}

func (service *ratingService) Update(ctx context.Context, dto *rating_dto.CreateDTO) error {
	var err error
	var rating *model.Rating

	if err := dto.Validate(); err != nil {
		return err
	}

	if dto.ItemId != nil {

		rating, err = service.findOneByItemId(ctx, dto)

	} else if dto.RestaurantId != nil {

		rating, err = service.findOneByRestaurantId(ctx, dto)
	}

	appErr, ok := err.(*common.AppError)

	isNotFound := ok && errors.Is(appErr.RootCauses(), gorm.ErrRecordNotFound)

	if isNotFound {

		if err := service.repo.Create(ctx, dto); err != nil {
			return common.ErrInternal(err).WithDebug(err.Error())
		}

	} else {

		if err := service.repo.Update(ctx, rating.Id, dto); err != nil {
			return common.ErrInternal(err).WithDebug(err.Error())
		}

		if dto.Like != nil {
			go func() {

				defer common.Recover()

				event := common.NewAppEvent(
					common.WithTopic(common.EventUserLikeRestaurant),
					common.WithData(dto.ToData()),
				)

				if err := service.eventPublisher.Publish(ctx, event.Topic, event); err != nil {
					log.Println("Failed to publish event", err)
				}
			}()
		}

	}

	return nil
}

func (service *ratingService) findOneByRestaurantId(ctx context.Context, dto *rating_dto.CreateDTO) (*model.Rating, error) {
	//there will have business logic before getting specific data with condition
	data, err := service.repo.FindOneWithCondition(ctx, map[string]any{
		"restaurant_id": dto.RestaurantId,
		"user_id":       dto.UserId,
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound(model.RatingEntity, err)
		}
		return nil, common.ErrInternal(err)
	}
	return data, nil
}

func (service *ratingService) findOneByItemId(ctx context.Context, dto *rating_dto.CreateDTO) (*model.Rating, error) {
	//there will have business logic before getting specific data with condition
	data, err := service.repo.FindOneWithCondition(ctx, map[string]any{
		"item_id": dto.ItemId,
		"user_id": dto.UserId,
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrEntityNotFound(model.RatingEntity, err)
		}
		return nil, common.ErrInternal(err)
	}
	return data, nil
}
