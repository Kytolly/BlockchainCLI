package block_database

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckDbEmpty(clt *mongo.Collection){
	var res bson.M
	err := clt.FindOne(context.TODO(), bson.D{}).Decode(&res)
	if err == mongo.ErrNoDocuments {
		slog.Debug("Collection is empty, insert genesis block...")
		clt.InsertOne(context.TODO(), pointerToLast_doc)
    }else if err != nil {
		slog.Error("Error query to collection!")
	}else {
		slog.Debug("Collection is not empty!")
	}
}