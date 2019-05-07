package model

import _"go.mongodb.org/mongo-driver/bson"


type User struct {
	Id        int           `json:"id" bosn:"id"`
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"password,omitempty" bson:"password"`
	Token     string        `json:"token,omitempty" bson:"-"`
}
