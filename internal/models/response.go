package models

type SuccessResponse struct {
	Success bool        `json:"success" default:"true"`
	Data    interface{} `json:"data"`
}

type FailureResponse struct {
	Success bool         `json:"success" default:"false"`
	Error   ErrorDetails `json:"error"`
}

type ErrorDetails struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
}
