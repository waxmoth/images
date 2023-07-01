package requests

// GetImage the request object for GetImage API
type GetImage struct {
	URL       string `form:"url" json:"url" binding:"required"`
	Type      string `form:"type" json:"type"`
	Width     int    `form:"width" json:"width"`
	Height    int    `form:"height" json:"height"`
	PositionX int    `form:"x" json:"x"`
	PositionY int    `form:"y" json:"y"`
}
