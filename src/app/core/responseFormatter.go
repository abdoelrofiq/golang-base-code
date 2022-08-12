package core

type Response struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func ResponseFormatter(message string, data interface{}) Response {
	return Response{
		Msg:  message,
		Data: data,
	}
}
