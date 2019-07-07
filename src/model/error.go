package model

import "github.com/vayan/sisistay/src/api/apiutils"

type ErrorResponse struct {
	Error string `json:"error"`
}

func SerializedErrorResponse(message string) []byte {
	return apiutils.Serialize(ErrorResponse{
		Error: message,
	})
}
