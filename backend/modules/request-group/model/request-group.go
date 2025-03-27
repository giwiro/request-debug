package model

import (
	"time"
)

type RequestGroup struct {
	Id        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Requests  []Request `json:"requests" bson:"requests"`
	CreateAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
