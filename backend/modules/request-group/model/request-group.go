package model

import (
	"time"
)

type RequestGroup struct {
	Id        string
	Requests  []Request
	CreateAt  time.Time
	UpdatedAt time.Time
}
