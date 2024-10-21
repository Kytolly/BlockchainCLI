package transaction_model

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log/slog"
	st"blockchain/pkg/setting"
)

var subsidy = st.Subsidy;

type Transaction struct{
	ID 		[]byte     // 交易的ID标识
	VIn 	[]TXInput  // 交易的输入
	VOut 	[]TXOutput // 交易的输出
}

func  NewCoinbaseTx(from, data string) *Transaction{
	// TODO:创建coinbase交易
	// 不需要以前存在的输出。它凭空创造了输出（即“硬币”）。
	if data == ""{
		data = fmt.Sprintf("Reward to '%s'", from) // 奖励给矿工
	}

	txin := TXInput{Txid:[]byte{}, Vout:-1, Signature: nil, PubKey: []byte(data)}
	txout := NewTXOutput(subsidy, from)
	
	tx := Transaction{ID:nil, VIn:[]TXInput{txin}, VOut:[]TXOutput{*txout}}
	tx.SetID()
	slog.Debug("coinbase built", "txid", fmt.Sprintf("%x", tx.ID))

	return &tx
}

func (tx *Transaction) SetID() {
	// TODO：为交易创建标识的ID
	tx.ID = tx.Hash()
}

// 判断交易是否为 coinbase 交易。
// 同时满足以下三个条件，说明是 coinbase 交易：
// 1. 只有一个输入；
// 2. 这个输入没有引用之前的交易；
// 3. 这个输入在之前的输出里索引为-1。
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.VIn) == 1 && len(tx.VIn[0].Txid) == 0 && tx.VIn[0].Vout == -1
}

func (tx *Transaction) Hash() []byte {
	//TODO： 获取交易的哈希值,为了保存在id中
	txCopy := *tx
	txCopy.ID = []byte{}
	hash := sha256.Sum256(txCopy.Serialize())
	return hash[:]
}


func (tx *Transaction) Serialize() []byte {
	// TODO： 序列化交易。
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}