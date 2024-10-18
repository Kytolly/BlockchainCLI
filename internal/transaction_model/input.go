package transaction_model

import(
	utils "blockchain/pkg/utils"
	"bytes"
)

type TXInput struct{
	Txid 		[]byte    // 输入的标识
	Vout    	int		  // 引用之前的Output，这个int是这个输出在交易输出序列中的索引
//	ScriptSig	string    // 验证签名需要对比公钥和签名数据
	Signature   []byte    // 签名，是一对数字的字节压缩
	PubKey      []byte    // 公钥（未哈希），是一对坐标的字节压缩
}

func(in *TXInput) UsesKey(pubKeyHash []byte) bool {
	// TODO：输入公钥的哈希值 是否和 引用输出的哈希值匹配
	// 确保发送者只花费属于他们的硬币
	lockingHash := utils.HashPubKey(in.PubKey) 
	return bytes.Equal(lockingHash, pubKeyHash)
}

// func(in *TXInput) CanUnlockOutputWith(unlockingData string) bool{
// 	return in.ScriptSig == unlockingData
// }