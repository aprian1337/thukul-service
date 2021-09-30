package responses

type TopUpResponse struct {
	Message string `json:"message"`
	Data    interface{}
}
