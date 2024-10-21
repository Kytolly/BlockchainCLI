package transaction_model

import (
	utils "blockchain/pkg/utils"
	"bytes"
	"encoding/gob"
)

// 引用输出的时候，输出将会作为一个整体被消费
type TXOutput struct{
	Value         int 		// 存储币的数量，单位为Satoshi
//	ScriptPubKtey string    // 用于解锁的公钥或脚本 
	PubKeyHash    []byte    // 输出的公钥哈希值
}

func NewTXOutput(value int, address string) *TXOutput {
	txo := &TXOutput{Value: value, PubKeyHash: nil}
	txo.Lock([]byte(address)) 
	return txo
}
func(out *TXOutput) Lock(address []byte) {
	// TODO: 通过钱包的地址，获得公钥哈希值，进行上锁
	// 向某人发送硬币时，我们只知道他们的地址 
	pubKeyHash := utils.GetPubKeyHashInAddress(address)
	out.PubKeyHash = pubKeyHash
}

func(out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	//TODO: 检查输出是否被锁定
	//匹配成功就可以进一步操作
	return bytes.Equal(out.PubKeyHash, pubKeyHash)
}

func (out *TXOutput) Serialize() []byte {
	// TODO： 序列化交易输出。
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(out)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}
// func(out *TXOutput) CanBeUnlockedWith(unlockingData string)bool{
// 	//TODO: 能否用address被解锁
// 	return out.ScriptPubKtey == unlockingData
// }