package models

import "labix.org/v2/mgo/bson"

const COLLECTION_NAME = "builds"
type build struct {
	id:int "id"
	title:string "title"
	build_category:category "category"
	version:string "version"
	updated:time "updated"
	created:time "created"
}
