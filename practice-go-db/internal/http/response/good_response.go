package response

type GoogResponse struct {
	Success bool `json:"success"`
}

func NewGoogResponse() *GoogResponse {
	return &GoogResponse{
		Success: true,
	}
}