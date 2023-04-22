package response

func BuildErrorResponse(message string, code int, data interface{}) Response {
	res := Response{
		Status:  false,
		Message: message,
		Code:    code,
		Data:    data,
	}

	return res
}
