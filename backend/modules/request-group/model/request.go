package model

import "time"

type RequestFile struct {
	Filename string `json:"filename" bson:"filename"`
	Size     int64  `json:"size" bson:"size"`
}

type Request struct {
	Id          string                   `json:"id,omitempty" bson:"id,omitempty"`
	Ip          string                   `json:"ip" bson:"ip"`
	Method      string                   `json:"method" bson:"method"`
	Host        string                   `json:"host" bson:"host"`
	BodySize    uint                     `json:"bodySize" bson:"bodySize"`
	BodyRaw     string                   `json:"bodyRaw" bson:"bodyRaw"`
	Url         string                   `json:"url" bson:"url"`
	Date        time.Time                `json:"date" bson:"date"`
	Files       map[string][]RequestFile `json:"files,omitempty" bson:"files,omitempty"`
	Form        map[string][]string      `json:"form,omitempty" bson:"form,omitempty"`
	QueryParams map[string]string        `json:"queryParams,omitempty" bson:"queryParams,omitempty"`
	Headers     map[string]string        `json:"headers,omitempty" bson:"headers,omitempty"`
}
