package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id   bson.ObjectId `jsoon:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}
