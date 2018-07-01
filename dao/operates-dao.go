package dao

import (
	. "github.com/xandeer/alpha-api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	OPERATES_COLLECTION = "operates"
)

// Find an operate by its id
func (m *DAO) FindOperatesGreaterThan(version int) ([]Operate, error) {
	var operates []Operate
	err := db.C(OPERATES_COLLECTION).Find(bson.M{"version": bson.M{"$gt": version}}).All(&operates)

	if err != nil {
		return nil, err
	}

	return operates, err
}

// Insert an operate into database
func (m *DAO) InsertOperate(operate Operate) error {
	err := db.C(OPERATES_COLLECTION).Insert(&operate)
	return err
}

func (m *DAO) GetLatestOperateVersion() (int, error) {
	version, err := db.C(OPERATES_COLLECTION).Find(bson.M{}).Count()

	if err == mgo.ErrNotFound {
		return 0, nil
	}

	return version, err
}
