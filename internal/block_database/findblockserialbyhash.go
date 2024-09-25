package block_database

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindBlockSerialByHash(clt *mongo.Collection, hash []byte) []byte{
	// TODO: 根据hash值在collection中查询并返回该区块的序列
	var doc document

	filter := bson.D{{Key:"key", Value:primitive.Binary{Subtype: 0x00, Data: hash}}}
	err := clt.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		slog.Error(err.Error()) 
        return nil
	}
	_, serial := readDocument(doc)
	return serial
}