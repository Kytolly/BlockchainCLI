package block_database

import (
	"context"
	"log/slog"
	//"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindLastBlock(clt *mongo.Collection, lastBlock interface{})bool{
	// TODO: 查询 MongoDB 数据库中不存在区块链->true
	filter := bson.D{primitive.E{Key: "timestamp", Value: bson.D{primitive.E{Key: "$exists", Value: true}}}}
	err := clt.FindOne(context.TODO(), filter).Decode(&lastBlock)
	if err == mongo.ErrNoDocuments {
        //fmt.Println(err)
		slog.Info("No block found in database...")
		return false
    }else if err != nil{
		slog.Error("Error finding last block!")
		slog.Error(err.Error())
		return false
	}

	slog.Info("The Last block found in database!")
	return true
    // 检查lastBlock 是否接收到对应的数据
}