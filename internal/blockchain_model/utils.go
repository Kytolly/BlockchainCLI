package blockchain_model

import (
	bm "blockchain/internal/block_model"
	"log/slog"
	"os"

	"github.com/boltdb/bolt"
)

func dbExists() bool {
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
    return nil
}

func(bc *BlockChain) GetBlock(id []byte) (block *bm.Block, err error) {
	return nil, nil
}