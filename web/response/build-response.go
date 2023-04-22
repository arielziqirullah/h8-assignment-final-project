package response

func BuildResponse(status bool, message string, code int, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Code:    code,
		Data:    data,
	}

	return res
}
