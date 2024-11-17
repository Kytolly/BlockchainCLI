package blockchain_model

import (
	bm "blockchain/internal/block_model"
	"errors"
	"log/slog"
	"os"

	"github.com/boltdb/bolt"
)

func dbExists(dbFile string) bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}

func(bc *BlockChain) GetBestHeight() int{
	// TODO: 返回最新的区块的高度
	var lastHeight int 
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lasthash := b.Get([]byte("l"))
		blockData := b.Get(lasthash)
		block := bm.DeserializeBlock(blockData)
		lastHeight = block.Height
		return nil
	})
	if err != nil {
		slog.Error(err.Error())
	}
	return lastHeight
}

func(bc *BlockChain) GetBlockHashes() [][]byte{
    // TODO：找到链上区块哈希的列表
	bci := bc.Iterator()
	var blocks [][]byte
	for{
		block := bci.Next()
		blocks = append(blocks, block.Hash)
		if len(block.PrevBlockHash)==0{
			break
		}
	}
	return blocks
}

func(bc *BlockChain) GetBlock(blockHash []byte) (block *bm.Block, err error) {
	// TODO: 根据哈希反过来查找区块
	err = bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket)) 
		blockData := b.Get(blockHash)
		if blockData == nil {
			return errors.New("the block may not be in the chain")
		}
		block = bm.DeserializeBlock(blockData)
		return nil
	})
	if err != nil {
		slog.Error(err.Error())
	}
	return block, nil
}