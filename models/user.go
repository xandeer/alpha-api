package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Represents an item, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Password string        `bson:"password" json:"password"`
}
