package requests

// GetImage the request object for GetImage API
type GetImage struct {
	URL    string `form:"url" json:"url" binding:"required"`
	Width  int16  `form:"width" json:"width"`
	Height int16  `form:"height" json:"height"`
}
