package block_database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func Inition()(*context.CancelFunc, *mongo.Collection, error){
	cancle, clt, err :=ConnectToDB()
	CheckDbEmpty(clt)
	return cancle, clt, err
}