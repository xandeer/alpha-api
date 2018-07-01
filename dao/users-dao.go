package dao

import (
	. "github.com/xandeer/alpha-api/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	USERS_COLLECTION = "users"
)

func (m *DAO) FindUserByName(name string) (User, error) {
	var user User
	err := db.C(USERS_COLLECTION).Find(bson.M{"name": name}).One(&user)

	return user, err
}
