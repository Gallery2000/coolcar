package objid

import (
	"coolcar/shared/id"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToAccountId(oid primitive.ObjectID) id.AccountID {
	return id.AccountID(oid.Hex())
}
