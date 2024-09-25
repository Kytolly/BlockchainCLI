package block_database

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertBlockToDb(db *mongo.Collection, hash []byte, serial []byte){
	// TODO:将创世区块插入collection

	doc := turnDocument(hash, serial)
	_, err := db.InsertOne(context.TODO(), doc)
	if err!= nil {
        slog.Error(err.Error())
        return
    }

	updatePointer(db, hash)
}

