package model

import "time"


type User struct {
	Id        string           `json:"id" bosn:"id"`
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"password,omitempty" bson:"password"`
	Token     string        `json:"token,omitempty" bson:"-"`
	Created   time.Time          `json:"created" bson:"created"`
}
