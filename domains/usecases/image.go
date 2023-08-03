package usecases

import (
	"context"
	"os"
	"time"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ImageService struct {
}

func NewImageService() *ImageService {
	return &ImageService{}
}

func (i *ImageService) GeneratePresignedUrl(ctx context.Context) (imageUploadresponse dto.ImageUploadResponse, err error) {
	// 파일 가져온 뒤, 버킷 업로드
	region := "ap-northeast-2"
	bucketName := os.Getenv("S3_BUCKET_NAME")
	accessKey := os.Getenv("AWS_ACCESS_KEY")
	secretKey := os.Getenv("AWS_SECRET_KEY")

	// AWS Config 생성
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		return
	}

	// Presigned URL 생성에 필요한 매개변수 설정
	objectKey := "hahaha.png" // S3에 저장될 파일 이름

	s3client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3client)
	presignedUrl, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	}, s3.WithPresignExpires(time.Minute*15))
	if err != nil {
		return
	}

	imageUploadresponse = dto.ImageUploadResponse{
		PresignedUrl: presignedUrl.URL,
	}

	return
}
