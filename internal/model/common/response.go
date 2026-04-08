package common

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) Success(data any) Response {
	return Response{
		Code:    200,
		Message: "OK",
		Data:    data,
	}
}

func (r *Response) Error(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
	}
}
