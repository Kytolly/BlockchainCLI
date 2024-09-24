package block_database

import (
	"context"
	"time"
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToDB() (*context.CancelFunc, *mongo.Collection, error){
	// TODO: 建立 MongoDB 连接并检查连接是否正常
	// 传递cancel保证能在main中关闭它
	uri := "mongodb://localhost:27017/"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
		return &cancel, nil, err
	}

	// 这里需要取消defer
	// 调用初始化函数后，数据库将被初始化，但一旦函数退出，数据库连接将会关闭
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()
	
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println(err)
		return &cancel, nil,err
	}

	clt := client.Database("MyBlockChain").Collection("block")

	slog.Info("The BlockChain is connected MongoDB successfully!")
	return &cancel, clt, nil
}
