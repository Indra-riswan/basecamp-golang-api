package helper

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type Emptyobj struct{}

func Succesresponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}

	return res
}

func Errorsresponse(message string, errors string, data interface{}) Response {
	Spliterr := strings.Split(errors, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  Spliterr,
		Data:    data,
	}

	return res
}
