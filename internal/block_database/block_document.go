package block_database

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

 

type document struct {
    Key   primitive.Binary `bson:"key"`
    Value primitive.Binary `bson:"value"`
}

var pointerToLast_doc = document{
    Key:   primitive.Binary{Subtype: 0x00, Data: []byte("last")},
    Value: primitive.Binary{Subtype: 0x00, Data: []byte{}},
}

func updatePointer(clt *mongo.Collection, hash []byte){
	filter := bson.D{{Key:"key", Value:primitive.Binary{Subtype: 0x00, Data: []byte("last")}}}
	target := bson.D{{Key:"value", Value: primitive.Binary{Subtype: 0x00, Data: hash}}}
	update := bson.D{{Key: "$set", Value: target}}
	_, err := clt.UpdateOne(context.TODO(), filter, update)
	if err!= nil{
		slog.Error(err.Error())
	}
}

// 将[]byte 转换成 document
func turnDocument(hash []byte, serial []byte) document {
	return document{
		Key:   primitive.Binary{Subtype: 0x00, Data: hash},
		Value: primitive.Binary{Subtype: 0x00, Data: serial},
	}
}

// 读取文档并还原为 []byte
func readDocument(result document) ([]byte, []byte) {
    return result.Key.Data, result.Value.Data 
}