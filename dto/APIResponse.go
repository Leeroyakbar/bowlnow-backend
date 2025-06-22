package dto

type APIResponse[T any] struct {
	Success      bool   `json:"success"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Data         T      `json:"data"`
	Message      string `json:"message"`
}
