package transaction_model

import(
	utils "blockchain/pkg/utils"
	st "blockchain/pkg/setting"
	"bytes"
)

// 引用输出的时候，输出将会作为一个整体被消费
type TXOutput struct{
	Value         int 		// 存储币的数量，单位为Satoshi
//	ScriptPubKtey string    // 用于解锁的公钥或脚本 
	PubKeyHash    []byte    // 输出的公钥哈希值
}

func(out *TXOutput) Lock(address []byte) {
	// TODO: 通过钱包的地址，获得公钥哈希值，进行上锁
	// 向某人发送硬币时，我们只知道他们的地址

	// 完整的未编码的地址包括版本号，公钥哈希值，校验和,需要去掉版本号和若干位校验和
	pubKeyHash := utils.Base58Decode(address)
	ChecksumLen := st.ChecksumLen
	out.PubKeyHash = pubKeyHash[1:len(pubKeyHash)-ChecksumLen] 
}

func(out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	//TODO: 检查输出是否被锁定
	//匹配成功就可以进一步操作
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// func(out *TXOutput) CanBeUnlockedWith(unlockingData string)bool{
// 	//TODO: 能否用address被解锁
// 	return out.ScriptPubKtey == unlockingData
// }