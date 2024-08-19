package responses

type UploadImage struct {
	Name string `form:"name" json:"name" example:"uploaded_image.png"`
}
