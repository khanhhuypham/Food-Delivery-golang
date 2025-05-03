package media_grpc_client

import (
	"Food-Delivery/pkg/common"
	"Food-Delivery/proto-buffer/gen/mediapb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type mediaGRPCClient struct {
	mediaGRPCServerURL string
	connect            *grpc.ClientConn
	client             mediapb.MediaClient
}

func NewMediaGRPCClient(mediaGRPCServerURL string) *mediaGRPCClient {
	connect, err := grpc.NewClient(
		mediaGRPCServerURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal(err)
	}
	client := mediapb.NewMediaClient(connect)

	return &mediaGRPCClient{
		mediaGRPCServerURL: mediaGRPCServerURL,
		connect:            connect,
		client:             client,
	}
}

func (grpc *mediaGRPCClient) UploadFiles(ctx context.Context, files []*mediapb.ImageUpload) ([]common.Image, error) {

	resp, err := grpc.client.UploadImages(ctx, &mediapb.UploadImagesRequest{Images: files})

	if err != nil {
		return nil, err
	}

	result := make([]common.Image, len(resp.Data))

	for i, file := range resp.Data {
		result[i] = common.Image{
			Id:     int(file.Id),
			Url:    file.Url,
			Size:   file.Size,
			Width:  int64ToIntPtr(file.Width),
			Height: int64ToIntPtr(file.Height),
		}
	}

	return result, nil
}

func int64ToIntPtr(v *int64) *int {
	if v == nil {
		return nil
	}
	val := int(*v)
	return &val
}
