package resources

type Response struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func JsonResponse(message string, data interface{}) Response {
	return Response{Message: message,Data: data}
}