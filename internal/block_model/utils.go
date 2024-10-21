package block_model

import (
	"bytes"
	"log/slog"
	// "crypto/sha256"
	tr "blockchain/internal/merkletree_model"
	"encoding/gob"
	"fmt"
)

// 区块相应的字段不应该明文存储进入数据库，将其转化为字节序列

func(b *Block) Serialize()[]byte{
	// TODO: 序列化区块，包含区块所有的信息
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil{
		fmt.Println(err)
	}
	return result.Bytes()
}

func DeserializeBlock(data []byte) *Block{
	// TODO: 反序列化区块，返回包含区块所有信息的结构体
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil{
		fmt.Println(err)
	}
	return &block
}

func(b *Block) HashTransactions()[]byte{
	// TODO: 计算交易序列打包的哈希值
	// var txHashes 	[][]byte
	// var txHash 		[32]byte
	var transactions [][]byte

	for _,tx := range b.Transactions{
		transactions = append(transactions, tx.Serialize())
	}
	//txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	slog.Debug("transactions", "len", len(transactions))
	mTree := tr.NewMerkleTree(transactions)
	return mTree.RootNode.Data
}