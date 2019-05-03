package model

import _"go.mongodb.org/mongo-driver/bson"


type User struct {
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"password,omitempty" bson:"password"`
	Token     string        `json:"token,omitempty" bson:"-"`
}
