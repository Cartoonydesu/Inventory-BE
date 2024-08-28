package response

type Response struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func SuccessResponse(st string, data any) (r Response){
	return Response {
		Status: st,
		Data: data,
	}
}

func ErrorResponse(st string, msg string) (r Response) {
	return Response{
		Status: st,
		Message: msg,
	}
}
