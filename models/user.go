package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id     bson.ObjectId // - Id in golang // json : "id" - postmand and  in bson : "_id" - mongodb
	Name   string        // - Name in golang // json : "name" - postmand and  in bson : "name" - mongodb
	Gender string        // - Gender in golang ....
	Age    int
}
