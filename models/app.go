package models

import "labix.org/v2/mgo/bson"

const COLLECTION_NAME = "apps"
type app struct {
	id:string "id"
	category:category "category"
	title:string "title"
	version:string "version"
	ip_addr:string "ip_addr"
	port:string "port"
	updated:time "updated"
	created:time "created"
}

