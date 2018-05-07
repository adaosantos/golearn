package models

import (

	mgo "gopkg.in/mgo.v2"
	"golearn/helpers/database"
)

// UserCollection returns the user collection
func UserCollection() *mgo.Collection {
	return database.Collection("users", mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}
