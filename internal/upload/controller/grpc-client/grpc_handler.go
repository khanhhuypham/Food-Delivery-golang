package media_grpc_client

import (
	"Food-Delivery/entity/model"
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

func (grpc *mediaGRPCClient) UploadFiles(ctx context.Context, files []*mediapb.ImageUpload) ([]model.Media, error) {

	resp, err := grpc.client.UploadImages(ctx, &mediapb.UploadImagesRequest{Images: files})

	if err != nil {
		return nil, err
	}

	result := make([]model.Media, len(resp.Data))

	for _, image := range resp.Data {
		log.Println(image)
	}

	return result, nil
}
