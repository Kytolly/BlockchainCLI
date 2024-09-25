package block_database

import (
	"context"
	"log/slog"

	//"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindLastBlockSerial(clt *mongo.Collection)[]byte{
	// TODO: 查询collection中的最后一个区块的序列
	var doc document
	var hash []byte

	filter := bson.D{{Key:"key", Value:primitive.Binary{Subtype: 0x00, Data: []byte("last")}}}
	err := clt.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		slog.Error(err.Error()) 
        return nil
	}

	_, hash = readDocument(doc)
	filter = bson.D{{Key:"key", Value:primitive.Binary{Subtype: 0x00, Data: hash}}}
	err = clt.FindOne(context.TODO(), filter).Decode(&doc)
	
	if err == mongo.ErrNoDocuments {
		slog.Info("No hash to block found in collection!")
		return nil
	}else if err != nil {
		slog.Error(err.Error())
		return nil
	}else {
		slog.Info("The Last block hash found in database, the serial of block back...")
		_, serial := readDocument(doc)
		return serial
	}
}