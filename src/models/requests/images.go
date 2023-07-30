package requests

import "mime/multipart"

// GetImage the request object for GetImage API
type GetImage struct {
	URL       string `form:"url" json:"url" binding:"required"`
	Type      string `form:"type" json:"type"`
	Width     int    `form:"width" json:"width"`
	Height    int    `form:"height" json:"height"`
	PositionX int    `form:"x" json:"x"`
	PositionY int    `form:"y" json:"y"`
}

// UploadImage the request object for UploadImage API
type UploadImage struct {
	Name string                `form:"name" json:"name"`
	File *multipart.FileHeader `form:"file" json:"file" binding:"required"`
}
