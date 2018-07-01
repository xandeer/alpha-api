package dao

import (
	. "github.com/xandeer/alpha-api/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	ITEMS_COLLECTION = "items"
)

// Find list of items
func (m *DAO) FindItemsAll() ([]Item, error) {
	var items []Item
	err := db.C(ITEMS_COLLECTION).Find(bson.M{}).All(&items)
	return items, err
}

// Find an item by its id
func (m *DAO) FindItemById(id string) (Item, error) {
	var item Item
	err := db.C(ITEMS_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&item)
	return item, err
}

// Insert an item into database
func (m *DAO) InsertItem(item Item) error {
	err := db.C(ITEMS_COLLECTION).Insert(&item)
	return err
}

// Delete an existing item
func (m *DAO) DeleteItemById(id string) error {
	err := db.C(ITEMS_COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

// Update an existing item
func (m *DAO) UpdateItem(item Item) error {
	err := db.C(ITEMS_COLLECTION).UpdateId(item.ID, &item)
	return err
}
