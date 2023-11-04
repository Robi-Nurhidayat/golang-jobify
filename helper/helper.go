package helper

type Response struct {
	Meta Meta
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
}

func ApiResponse(msg string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: msg,
		Status:  status,
		Code:    code,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}