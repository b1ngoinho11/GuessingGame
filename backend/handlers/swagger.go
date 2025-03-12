package handlers

type LoginRequest struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"admin"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type GuessResponse struct {
	Message string `json:"message"`
}
