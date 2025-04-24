package upload

import (
	"Food-Delivery/config"
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"Food-Delivery/proto-buffer/gen/mediapb"
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type UploadProvider interface {
	UploadFile(ctx *gin.Context) (*model.Media, error)
	UploadFiles(ctx context.Context, files []*mediapb.ImageUpload) ([]*model.Media, error)
	DeleteFile(ctx context.Context, destination string) error
}

type s3Provider struct {
	bucket  string
	region  string
	apiKey  string
	secret  string
	domain  string
	session *session.Session
}

func NewS3Provider(cfg *config.Config) *s3Provider {
	provider := &s3Provider{
		bucket: cfg.Aws.S3Bucket,
		region: cfg.Aws.Region,
		apiKey: cfg.Aws.APIKey,
		secret: cfg.Aws.SecretKey,
		domain: cfg.Aws.S3Domain,
	}
	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(
			provider.apiKey, //Access key ID
			provider.secret, // secret access key
			"",              //Token co thể bỏ qua
		),
	})

	if err != nil {
		log.Fatalln(err)
	}
	provider.session = s3Session
	return provider
}

func (provider *s3Provider) UploadFile(ctx *gin.Context) (*model.Media, error) {
	fileHeader, err := ctx.FormFile("files")
	if err != nil {
		panic(common.ErrBadRequest(err))
	}

	Id, _ := uuid.NewV7()
	folder := ctx.DefaultPostForm("folder", "photo")
	fileName := fmt.Sprintf("%s%s", Id, filepath.Ext(fileHeader.Filename))
	file, err := fileHeader.Open()
	if err != nil {
		panic(common.ErrBadRequest(err))
	}
	defer file.Close()

	// Read file content into bytes
	dataBytes, err := io.ReadAll(file)
	if err != nil {
		panic(common.ErrBadRequest(err))
	}

	// Get content type
	contentType := http.DetectContentType(dataBytes)

	//Lấy width, height của image
	width, height, err := getImageDimension(dataBytes)
	if err != nil {
		//File không phải là hình ảnh
		panic(common.ErrBadRequest(err))
	}

	destination := fmt.Sprintf("%s/%s", folder, fileName)

	// Upload to S3
	_, err = s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(provider.bucket),
		Key:         aws.String(destination),
		ACL:         aws.String("private"),
		ContentType: aws.String(contentType),
		Body:        bytes.NewReader(dataBytes),
	})

	if err != nil {
		return nil, err
	}

	img := &model.Media{
		Folder:    folder,
		Filename:  fileName,
		CloudName: "aws-s3",
		Url:       fmt.Sprintf("%s/%s", provider.domain, destination),
		Size:      fileHeader.Size,
		Height:    &height,
		Width:     &width,
		Ext:       strings.ReplaceAll(filepath.Ext(fileName), ".", ""),
	}
	return img, nil
}

func (provider *s3Provider) UploadFiles(ctx context.Context, files []*mediapb.ImageUpload) ([]*model.Media, error) {
	if len(files) == 0 {
		return nil, nil
	}

	var results []*model.Media
	folder := "photo"

	for _, file := range files {
		Id, _ := uuid.NewV7()
		fileName := fmt.Sprintf("%s%s", Id, filepath.Ext(file.Filename))

		// Lấy width, height của image
		width, height, err := getImageDimension(file.Content)
		if err != nil {
			// File không phải là hình ảnh
			return nil, common.ErrBadRequest(err)
		}

		destination := fmt.Sprintf("%s/%s", folder, fileName)

		// Upload to S3
		_, err = s3.New(provider.session).PutObject(&s3.PutObjectInput{
			Bucket:      aws.String(provider.bucket),
			Key:         aws.String(destination),
			ACL:         aws.String("private"),
			ContentType: aws.String(file.ContentType),
			Body:        bytes.NewReader(file.Content),
		})

		if err != nil {
			return nil, err
		}

		img := &model.Media{
			Folder:    folder,
			Filename:  fileName,
			CloudName: "aws-s3",
			Url:       fmt.Sprintf("%s/%s", provider.domain, destination),
			Height:    &height,
			Width:     &width,
			Ext:       strings.ReplaceAll(filepath.Ext(fileName), ".", ""),
		}

		results = append(results, img)
	}

	return results, nil
}

//func (provider *s3Provider) UploadFiles(ctx context.Context, files []*mediapb.ImageUpload) ([]*model.Media, error) {
//	if len(files) == 0 {
//		return nil, nil
//	}
//
//	var wg sync.WaitGroup
//	var mu sync.Mutex
//	var results []*model.Media
//	folder := "photo"
//
//	for _, file := range files {
//		wg.Add(1)
//		go func(file *mediapb.ImageUpload) {
//			defer wg.Done()
//
//			Id, _ := uuid.NewV7()
//			fileName := fmt.Sprintf("%s%s", Id, filepath.Ext(file.Filename))
//
//			width, height, err := getImageDimension(file.Content)
//			if err != nil {
//				return // skip this file, or send error to a channel if you want to collect them
//			}
//
//			destination := fmt.Sprintf("%s/%s", folder, fileName)
//
//			_, err = s3.New(provider.session).PutObject(&s3.PutObjectInput{
//				Bucket:      aws.String(provider.bucket),
//				Key:         aws.String(destination),
//				ACL:         aws.String("private"),
//				ContentType: aws.String(file.ContentType),
//				Body:        bytes.NewReader(file.Content),
//			})
//
//			if err != nil {
//				return
//			}
//
//			img := &model.Media{
//				Folder:    folder,
//				Filename:  file.Filename,
//				CloudName: "aws-s3",
//				Url:       fmt.Sprintf("%s/%s", provider.domain, destination),
//				Height:    &height,
//				Width:     &width,
//				Ext:       strings.ReplaceAll(filepath.Ext(fileName), ".", ""),
//			}
//
//			mu.Lock()
//			results = append(results, img)
//			mu.Unlock()
//		}(file)
//	}
//
//	wg.Wait()
//	return results, nil
//}

/*
Comparison: Sequential vs Concurrent UploadFiles

Feature / Concern          | Sequential Version (Single-threaded)     | Concurrent Version (with goroutines)
-----------------------------------------------------------------------------------------------
Concurrency                | No — uploads happen one-by-one           | Yes — uploads happen in parallel
Performance                | Slower for large file lists              | Faster with parallel uploads
Resource Usage             | Very predictable                         | Higher memory and network usage if unbounded
Error Handling             | Stops immediately on the first error     | Skips failed uploads silently (can be improved)
Thread Safety              | No concern needed                        | Requires sync.Mutex to avoid race conditions
Result Order (Determinism) | Maintains order of input files           | May return results in different order
Code Complexity            | Simple, easy to reason about             | More complex: uses sync.WaitGroup + Mutex
*/

func (provider *s3Provider) DeleteFile(ctx context.Context, destination string) error {

	// delete file from S3

	_, err := s3.New(provider.session).DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(provider.bucket),
		Key:    aws.String(destination),
	})

	if err != nil {
		return err
	}
	return nil
}

func getImageDimension(dataBytes []byte) (int, int, error) {
	fileBytes := bytes.NewBuffer(dataBytes)
	img, _, err := image.DecodeConfig(fileBytes)
	if err != nil {
		return 0, 0, nil
	}
	return img.Width, img.Height, nil
}
