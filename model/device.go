package model

import "time"

type App struct {
	Name string `json:"name" bson:"name"`
	Version string 'json: "version" bson:"version"'
	UpdateDate time.Time 'json:"updatedate" bson:"updatedate"'
}

type Device struct {
	Id        string        `json:"id" bson:"id"`
	Name     string        `json:"name" bson:"name"`
	Ipv4      string        `json:"ipv4" bson:"ipv4"`
	VncPort      string        `json:"vnc_port" bson:"vnc_port"`
	Apps []app `json:"apps" bson:"apps"`
	Updated   time.Time           `json:"update" bson:"update"` 
	Created   time.Time          `json:"created" bson:"created"` 
}
