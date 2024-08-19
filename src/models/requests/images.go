package requests

import "mime/multipart"

// GetImage the request object for GetImage API
type GetImage struct {
	URL       string `form:"url" json:"url" binding:"required" example:"https://example.com/image.png"`
	Type      string `form:"type" json:"type" enums:"crop,resize" example:"crop" default:"crop"`
	Width     int    `form:"width" json:"width" example:"100"`
	Height    int    `form:"height" json:"height" example:"100"`
	PositionX int    `form:"x" json:"x"`
	PositionY int    `form:"y" json:"y"`
}

// UploadImage the request object for UploadImage API
type UploadImage struct {
	Name string                `form:"name" json:"name" example:"image.png"`
	File *multipart.FileHeader `form:"file" json:"file" binding:"required" swaggerignore:"true"`
}
