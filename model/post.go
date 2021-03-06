package model

import "time"

type Post struct {
	Id        string           `json:"id" bosn:"id"`
	Title string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
	Writer string `json:"writer" bson:"writer"`
	Updated time.Time `json:"update" bson:"update"`
	Created time.Time `json:"created" bson:"created"`
}


