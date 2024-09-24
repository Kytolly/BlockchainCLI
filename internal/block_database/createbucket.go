package block_database

import (
	"context"
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
)

// bm "blockchain/internal/block_model"
//"context"
//"fmt"
func CreateBucket(db *mongo.Collection, genesis interface{}){
	// TODO:将创世区块插入数据库

	slog.Debug("createBucket:", "genesis", fmt.Sprintf("%v", genesis))
	_, err := db.InsertOne(context.TODO(), genesis)

	slog.Debug("createBucket:", "err", fmt.Sprintf("%v", err))
	if err != nil {
		slog.Error("Create blockchain Bucket failed!")
	}
	slog.Info("Blockchain bucket created successfully!")
}

