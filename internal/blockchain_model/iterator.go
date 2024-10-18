package blockchain_model

import ( 
	bm "blockchain/internal/block_model"
	"log/slog"

	"github.com/boltdb/bolt" 
)

type BlockChainIterator struct{
	currentHash []byte
	db 			*bolt.DB
}

func(bc BlockChain) Iterator()*BlockChainIterator{
	// TODO：初始化区块链迭代器
	// 迭代器最初指向区块链的顶端，从新到旧地获取区块
	bci := &BlockChainIterator{bc.tip, bc.db}
	return bci
}

func(i *BlockChainIterator) Next() *bm.Block{
	// TODO: 迭代器的前一个区块，更新迭代器
	var block *bm.Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = bm.DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		slog.Error(err.Error())
	}

	i.currentHash = block.PrevBlockHash
	return block
}

