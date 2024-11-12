package wallet_model

import (
	"crypto/ecdsa" 
	// st "blockchain/pkg/setting"
	utils "blockchain/pkg/utils"
)

// var version = []byte(st.AlgorithmVersion)
var version = []byte{0x00}

type Wallet struct {
	// 用私钥和公钥代表钱包所有者的身份
	PrivateKey ecdsa.PrivateKey  // 私钥不应该被暴露
	PublicKey []byte  
}

func NewWallet() *Wallet {
	//TODO: 初始化钱包
	private, public := newKeyPair()
	wallet := Wallet{PrivateKey: private, PublicKey: public}
	return &wallet
}

func(w *Wallet) GetAddress()[]byte {
	//TODO： 根据钱包生成地址

	pubKeyHash := utils.HashPubKey(w.PublicKey)

	// 地址生成算法的版本号添加在hash值前面
	versionedPayload := append(version, pubKeyHash...)

	checksum := checksum(versionedPayload)

	// 完整的地址包括版本号，公钥哈希值，校验和，用Base58编码
	fullPayload := append(versionedPayload, checksum...)
	address := utils.Base58Encode(fullPayload)

	return address
}

