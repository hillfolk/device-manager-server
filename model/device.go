package model

import "time"

type Device struct {
	Id        string        `json:"id" bosn:"id"`
	Title     string        `json:"title" bson:"title"`
	Ipv4      string        `json:"ipv4" bson:"ipv4"`
	Port      string        `json:"port" bson:"port"`
	Updated   time.Time           `json:"update" bson:"update"` 
	Created   time.Time          `json:"created" bson:"created"` 
}
