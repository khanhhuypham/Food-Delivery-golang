package media_service

import (
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"Food-Delivery/pkg/upload"
	"context"
	"fmt"
)

type MediaRepository interface {
	Create(ctx context.Context, media *model.Media) error
	FindDataWithCondition(ctx context.Context, condition map[string]any) (*model.Media, error)
	Delete(ctx context.Context, condition map[string]any) error
}

type MediaService interface {
	DeleteFile(ctx context.Context, destination string) error
}

type mediaService struct {
	mediaRepo  MediaRepository
	s3Provider upload.UploadProvider
}

func NewMediaService(mediaRepo MediaRepository, s3Provider upload.UploadProvider) *mediaService {
	return &mediaService{
		mediaRepo,
		s3Provider,
	}
}

func (service *mediaService) Create(ctx context.Context, media *model.Media) error {

	if err := service.mediaRepo.Create(ctx, media); err != nil {
		return common.ErrInternal(err).WithDebug(err.Error())
	}
	return nil
}

func (service *mediaService) Delete(ctx context.Context, id int) {

	media := service.FindOneById(ctx, id)

	if media != nil {

		if err := service.mediaRepo.Delete(ctx, map[string]any{"id": id}); err != nil {
			return
		}

		if media.Folder != "" {
			fileName := fmt.Sprintf("%s/%s", media.Folder, media.Filename)
			if err := service.s3Provider.DeleteFile(ctx, fileName); err != nil {
				return
			}
		} else {
			if err := service.s3Provider.DeleteFile(ctx, media.Filename); err != nil {
				return
			}
		}

	}

}

func (service *mediaService) FindOneById(ctx context.Context, id int) *model.Media {

	media, err := service.mediaRepo.FindDataWithCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return nil
	}
	return media
}
