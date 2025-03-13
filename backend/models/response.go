package models

type LoginResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type GuessResponse struct {
	Message string `json:"message"`
}
