package model

import "time"

type App struct {
	Name string `json:"name" bson:"name"`
	Version string `json:"version" bson:"version"`
	LastUpdate time.Time `json:"lastUpdate" bson:"lastUpdate"`
	TargetVersion string  `json:"target_version" bson:"target_version"`
	UpdateRequest bool  `json:"update_requst" bson:"update_request"`
	
}

type Device struct {
	Id        string        `json:"id" bson:"id"`
	Name     string        `json:"name" bson:"name"`
	Ipv4      string        `json:"ipv4" bson:"ipv4"`
	VncPort      string        `json:"vnc_port" bson:"vnc_port"`
	Apps []App `json:"apps" bson:"apps"`
	Owner string `json:"owner" bson:"owner"` 
	Updated   time.Time           `json:"update" bson:"update"` 
	Created   time.Time          `json:"created" bson:"created"` 
}
