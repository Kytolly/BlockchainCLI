package block_model

import (
	bdb "blockchain/internal/block_database"
	ts "blockchain/internal/transaction_model"

	"context"
	"fmt"
	"log/slog"


	"go.mongodb.org/mongo-driver/mongo"
)

var genesisCoinbaseData = "Genesis Block!"


type BlockChain struct{
	// blocks []*Block
	tip []byte
	Db  *mongo.Collection
}

func (bc *BlockChain) MineBlock(transactions []*ts.Transaction) {
	//TODO: 向区块链添加新区块
	var lasthash []byte
	// 验证发生在交易进入区块之前
	for _, tx := range transactions {
		if !bc.VerifyTransaction(tx) {
			slog.Warn("Waning: Invalid transaction!!!")
		}
	}
	lastblock_serial := bdb.FindLastBlockSerial(bc.Db)
	lastBlock := DeserializeBlock(lastblock_serial)
	lasthash = lastBlock.Hash

	// newBlock := NewBlock(data, lasthash)
	newBlock := NewBlock(transactions, lasthash)
	newBlock_serial := newBlock.Serialize()	
	newBlock_hash := newBlock.Hash[:]

	bdb.InsertBlockToDb(bc.Db, newBlock_hash, newBlock_serial)
	bc.tip = newBlock.Hash
}

func NewBlockChain(address string) (*context.CancelFunc, *BlockChain) {
	//TODO: 创建一个新的区块链
	//地址将获得挖掘创世区块的奖励。
    // return &BlockChain{blocks: []*Block{NewGenesisBlock()}}

	// 打开一个DB
	var tip []byte
	cancle , db, _ := bdb.Inition()

	// 检查db中是否存储了区块链
	lastblock_serial := bdb.FindLastBlockSerial(db)
	if lastblock_serial != nil{
		// 如果存在，上一次运行程序会把tip保存为doc
		// 读取对应的doc，设置尖端为最后一个区块的Hash(上一次tip)
		lastBlock := DeserializeBlock(lastblock_serial)
		tip = lastBlock.Hash
		slog.Info("The tip is be set to the last block hash...")
		
		return cancle,&BlockChain{tip, db}
	}else {
		// 如果不存在，创建创世区块，创建存储桶
		// 将区块保存其中，更新最后一个区块的Hash密钥
		// 注意存入区块的是包含区块所有信息的一个字节序列

		// 创建创世区块之前，先创建egg：一个coinbase交易
		cbtx := ts.NewCoinbaseTx(address, genesisCoinbaseData)
        genesis := NewGenesisBlock(cbtx)
		slog.Debug("coinbase TX in address", "", address)

		genesis_serial := genesis.Serialize()
		genesis_hash := genesis.Hash[:]

        bdb.InsertBlockToDb(db, genesis_hash, genesis_serial)
		tip = genesis_hash
		slog.Debug("NewBlockChain:", "tip", fmt.Sprintf("%x", tip))
	}
	fmt.Printf("tip: %x\n", tip)
	fmt.Println("New BlockChain has been build!")
	return cancle, &BlockChain{tip, db}
}


