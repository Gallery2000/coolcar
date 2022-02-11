package dao

import (
	"context"
	"coolcar/shared/id"
	mgutil "coolcar/shared/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const openIDField = "open_id"

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}

func (m *Mongo) ResolveAccountId(c context.Context, openID string) (id.AccountID, error) {
	inserteID := mgutil.NewObjID()
	res := m.col.FindOneAndUpdate(c, bson.M{
		openIDField: openID,
	}, mgutil.SetOnInsert(bson.M{}))
}
