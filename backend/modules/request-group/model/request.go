package model

import "time"

type Request struct {
	Id          string            `json:"id,omitempty" bson:"id,omitempty"`
	Ip          string            `json:"ip" bson:"ip"`
	Method      string            `json:"method" bson:"method"`
	Host        string            `json:"host" bson:"host"`
	Date        time.Time         `json:"date" bson:"date"`
	QueryParams map[string]string `json:"queryParams,omitempty" bson:"queryParams,omitempty"`
}
