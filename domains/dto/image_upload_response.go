package dto

type ImageUploadResponse struct {
	PresignedUrl string `json:"presigned_url"`
	ImageUrl     string `json:"image_url"`
}
