package custom_error

type CustomError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (custom CustomError) Error() string {
	return custom.Message
}


