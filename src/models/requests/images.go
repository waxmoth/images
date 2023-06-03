package requests

type GetImage struct {
	Url    string `form:"url" json:"url" binding:"required"`
	Width  int16  `form:"width" json:"width"`
	Height int16  `form:"height" json:"height"`
	Name   string `form:"name" json:"name"`
}
