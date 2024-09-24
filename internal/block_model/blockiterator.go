package block_model

import (
	//"context"
	"fmt"
	bdb "blockchain/internal/block_database"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlockChainIterator struct{
	currentHash []byte
	db 			*mongo.Collection
}

func(bc BlockChain) Iterator()*BlockChainIterator{
	// TODO：初始化区块链迭代器
	// 迭代器最初指向区块链的顶端，从新到旧地获取区块
	bci := &BlockChainIterator{bc.tip, bc.Db}
	return bci
}

func(i *BlockChainIterator) Next() *Block{
	// TODO: 迭代器的下一个区块，更新迭代器
	var block *Block

	if bdb.FindLastBlock(i.db, &block){
		fmt.Printf("No Last Block Found!")
		return nil
	}
	i.currentHash = block.PrevBlockHash
	return block
}