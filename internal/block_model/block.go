package block_model

import (
	ts "blockchain/internal/transaction_model" 
	"time"
)

type BlockHeader struct {
	Timestamp      int64  	// 创建区块时的当前时间戳
	Nonce          int		// 随机数
    PrevBlockHash  []byte 	// 前一个区块的哈希值
//  Version        string 	// 区块版本
//  MerkleRoot     string 	// 包含所有交易数据的默克根
//  Difficulty     int    	// 工作量证明难度
}

type Block struct{
	BlockHeader 			
	//Data 			[]byte 	// 区块的交易数据(原始)
	Transactions	[]*ts.Transaction //区块的交易数据
	Hash  			[]byte 	// 区块的哈希值
}

func (b *Block) SetHash() {
	// TODO: 计算区块的哈希值

	// 进行字符串的拼接获得Hash值，这样的区块简单又便宜
	// timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	// headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	// hash := sha256.Sum256(headers)

	// 需要完成工作量证明算法，生成哈希值
	pow := NewProofOfWork(b)
	nonce, hash := pow.Run()

	b.Nonce = nonce
	b.Hash = hash[:]
}

func NewBlock(transactions []*ts.Transaction, prevBlockHash []byte) *Block {
	//TODO: 创建一个新的区块
	block := &Block{
		//Timestamp: time.Now().Unix(),
		//PrevBlockHash: prevBlockHash,
		BlockHeader: BlockHeader{
			Timestamp:  	time.Now().Unix(),
			Nonce: 	    	0,
			PrevBlockHash:  prevBlockHash,
		//  Version:        "0.1.1",
		//  MerkleRoot:     "",
        //  Difficulty:     10,
        },
        //Data: []byte(data),
		Transactions: transactions,
		Hash: []byte{},
	}
	block.SetHash()
	return block
}

func NewGenesisBlock(coinbase *ts.Transaction) *Block{
	//TODO: 建立创世块
	//return NewBlock("Genesis Block!", []byte{})
	return NewBlock([]*ts.Transaction{coinbase}, []byte{})
}