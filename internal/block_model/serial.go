package block_model

import(
	"bytes"
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