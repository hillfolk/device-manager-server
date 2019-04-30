package models

import "labix.org/v2/mgo/bson"

const COLLECTION_NAME = "categorys"
type category struct {
	id:int "id"
	title:string "title"
	created:time "created"
}
