package models

import "labix.org/v2/mgo/bson"

const COLLECTION_NAME = "posts"

type post struct{
	id:int "id"
	title:string "title"
	content:string "content"
	user:string "user"
	updated:time "updated"
	created:time	"created
}
