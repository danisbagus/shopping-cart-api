package common

type DefaultResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const TIME_FORMAT = "2006-01-02 15:04:05"

func (r *DefaultResponse) SetResponseData(message string, data interface{}) {
	r.Message = message
	r.Data = data
}

func (r *DefaultResponse) SetResponseWithoutData(message string) {
	r.Message = message
}
