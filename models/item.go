package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Represents an item, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Item struct {
	ID       bson.ObjectId `bson:"_id" json:"_id"`
	Content  string        `bson:"content" json:"content"`
	Author   string        `bson:"author" json:"author"`
	From     string        `bson:"from" json:"from"`
	Tags     []string      `bson:"tags" json:"tags"`
	Hash     string        `bson:"hash" json:"hash"`
	Removed  bool          `bson:"removed" json:"removed"`
	Created  time.Time     `bson:"created" json:"created"`
	Modified time.Time     `bson:"modified" json:"modified"`
}
