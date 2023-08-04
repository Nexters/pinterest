package usecases

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type ImageService struct {
}

func NewImageService() *ImageService {
	return &ImageService{}
}

func (i *ImageService) GeneratePresignedUrl(ctx context.Context, filename string) (imageUploadresponse dto.ImageUploadResponse, err error) {
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

	imageName := uuid.New().String()
	ext := getExtension(filename)
	objectKey := fmt.Sprintf("%s.%s", imageName, ext) // S3에 저장될 파일 이름

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

func getExtension(image string) string {
	imageName := strings.ToLower(image)

	// 파일 이름에서 마지막 점 이후의 부분을 확장자로 간주
	parts := strings.Split(imageName, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}

	// 확장자가 없을 경우 jpg로 기본 설정
	return "jpg"
}
