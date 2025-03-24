package model

import "time"

type ApiError struct {
	Error     string        `json:"error"`
	Message   string        `json:"message"`
	Status    int           `json:"status"`
	SubErrors []ApiSubError `json:"subErrors,omitempty"`
	Timestamp time.Time     `json:"timestamp"`
}
