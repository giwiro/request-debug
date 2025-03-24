package model

type ApiSubError struct {
	Error   string `json:"error"`
	Field   string `json:"field"`
	Message string `json:"message"`
}
