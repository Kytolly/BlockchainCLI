package blockchain_model

import (
	bm "blockchain/internal/block_model"
	ts "blockchain/internal/transaction_model"
	st "blockchain/pkg/setting"
	"fmt"
	"log/slog"
	"os"

	"github.com/boltdb/bolt"
)

var genesisCoinbaseData = st.GenesisCoinbaseData
var dbFile = st.ChainDbFile
var blocksBucket = st.BlockBucket

type BlockChain struct{
	// blocks []*Block
	tip []byte
	db  *bolt.DB
}

func (bc *BlockChain) MineBlock(transactions []*ts.Transaction)*bm.Block {
	//TODO: 向区块链添加新区块
	var lasthash []byte
	// 验证发生在交易进入区块之前
	for _, tx := range transactions {
		if !bc.VerifyTransaction(tx) {
			slog.Warn("Waning: Invalid transaction!!!")
		}
	}
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lasthash = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		slog.Error(err.Error())
	}

	newBlock := bm.NewBlock(transactions, lasthash)
	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			slog.Error(err.Error())
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			slog.Error(err.Error())
		}

		bc.tip = newBlock.Hash

		return nil
	})
	if err != nil {
		slog.Error(err.Error())
	}
	bc.tip = newBlock.Hash
	slog.Info("A Block is mined successfully!")
	fmt.Printf("A Block is mined successfully!\n")
	return newBlock
}

func NewBlockChain(address string) *BlockChain {
	//TODO: 创建一个新的区块链
	if dbExists(){
		slog.Info("BlockChain already exists")
		os.Exit(1) 
	}

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil) 
	if err != nil {
		slog.Warn("Failed to open chain dbfile")
		return nil
	}

	cbtx := ts.NewCoinbaseTx(address, genesisCoinbaseData)
	genesis := bm.NewGenesisBlock(cbtx)

	err = db.Update(func(tx *bolt.Tx) error{
		b, err := tx.CreateBucket([]byte(blocksBucket))
		if err != nil {
			slog.Error(err.Error())
			return err
		}
		err = b.Put(genesis.Hash, genesis.Serialize())
		if err != nil {
			slog.Error(err.Error())
			return err
		}
		
		err = b.Put([]byte("l"), genesis.Hash)
		if err!= nil {
            slog.Error(err.Error())
            return err
        }

		tip = genesis.Hash
		return nil
	})
	if err != nil {
		slog.Error(err.Error())
        return nil
	}
	bc := BlockChain{tip, db}
	return &bc
}

func GetBlockChain() *BlockChain {
	//TODO: 利用已创建的区块链
	if !dbExists() {
		slog.Info("BlockChain not exists")
		return nil 
	}
	var tip []byte
	db, err := bolt.Open(dbFile, 0000, nil) 
	if err != nil {
		slog.Info(err.Error())
		return nil
	}
	err = db.Update(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(blocksBucket))
        tip = b.Get([]byte("l"))
        return nil
	})
	if err != nil {
		slog.Error(err.Error())
		return nil
	}
	bc := BlockChain{tip, db}
	return &bc
}

func(bc *BlockChain) Close(){
	bc.db.Close()
}