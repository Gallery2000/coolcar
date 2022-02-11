package mgutil

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	IDFieldName       = "_id"
	updateAtFieldName = "updateat"
)

var NewObjID = primitive.NewObjectID

func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}
