package mobile

// Standard API response structure for mobile endpoints
type ApiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
}