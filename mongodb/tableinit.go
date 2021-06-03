package mongodb

import (
	"gopkg.in/mgo.v2"
)

var (
	tb string = "Blocks"
	tt string = "Transactions"
	collectionBlocks *mgo.Collection
	collectionTransactions *mgo.Collection
)

// do this when reconnect to the database
func deinintCollections() {
	collectionBlocks = database.C(tb)
	collectionTransactions = database.C(tt)
}

