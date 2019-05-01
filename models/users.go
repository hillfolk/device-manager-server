package models
import "labix.org/v2/mgo/bson"

const COLLECTION_NAME = "users"

type user struct {
	id:int "id"
	name:string "name"
	password:string "password"
	created:time "time"
}
