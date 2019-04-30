package models

import "labix.org/v2/mgo/bson"

const COLLECTION_NAME = "categorys"

type post struct{
	id:int
	title:string
	content:string
	user:string
	updated:time
	created:time	
}
