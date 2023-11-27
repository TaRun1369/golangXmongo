package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    Id     primitive.ObjectID `bson:"_id,omitempty"` // Id in golang // json : "id" - postman and  in bson : "_id" - mongodb
    Name   string             `bson:"name"`          // Name in golang // json : "name" - postman and  in bson : "name" - mongodb
    Gender string             `bson:"gender"`        // Gender in golang ....
    Age    int                `bson:"age"`
}