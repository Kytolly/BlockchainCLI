package block_model

import (
	//"bytes"
	//"crypto/sha256"
	//"strconv"
	bdb "blockchain/internal/block_database"
	
	"context"

	"fmt"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
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
	BlockHeader 			`bson:",inline"`
	Data 			[]byte 	`bson:"data"` // 区块的交易数据
	Hash  			[]byte 	`bson:"hash"` // 区块的哈希值
}

type BlockChain struct{
	// blocks []*Block
	tip []byte
	Db  *mongo.Collection
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

func NewBlock(data string, prevBlockHash []byte) *Block {
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
        Data: []byte(data),
		Hash: []byte{},
	}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block{
	//TODO: 建立创世块
	return NewBlock("Genesis Block!", []byte{})
}

func (bc *BlockChain) AddBlock(data string) {
	//TODO: 向区块链添加新区块
    // prevBlock := bc.blocks[len(bc.blocks)-1]
    // newBlock := NewBlock(data, prevBlock.Hash)
    // bc.blocks = append(bc.blocks, newBlock)

	var lasthash []byte
	lastblock_serial := bdb.FindLastBlockSerial(bc.Db)
	lastBlock := DeserializeBlock(lastblock_serial)
	lasthash = lastBlock.Hash

	newBlock := NewBlock(data, lasthash)
	newBlock_serial := newBlock.Serialize()	
	newBlock_hash := newBlock.Hash[:]

	bdb.InsertBlockToDb(bc.Db, newBlock_hash, newBlock_serial)
	bc.tip = newBlock.Hash
}

func NewBlockChain() (*context.CancelFunc, *BlockChain) {
    //TODO: 创建一个新的区块链
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
        genesis := NewGenesisBlock()
		genesis_serial := genesis.Serialize()
		genesis_hash := genesis.Hash[:]

        bdb.InsertBlockToDb(db, genesis_hash, genesis_serial)
		tip = genesis_hash
		slog.Debug("NewBlockChain:", "tip", fmt.Sprintf("%x", tip))
	}
	return cancle, &BlockChain{tip, db}
}