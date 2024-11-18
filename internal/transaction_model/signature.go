package transaction_model

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log/slog"
	"math/big"
)

func(tx *Transaction) Sign(privKey ecdsa.PrivateKey, prevTXs map[string]Transaction){
	// TODO: 私钥签名
	// 交易会解锁以前的输出，重新分配它们的值，并锁定新的输出
	// 必须对交易的发送者，接受者，新输出的值进行签名，需要一些修剪
	// 哈希公钥值存储在未锁定的输出中，和新的锁定的输出中

	// coinbase交易无需签名，因为coinbase 交易中没有实际输入
	if tx.IsCoinbase() {
		return 
	}
	for _, vin := range tx.VIn {
		if prevTXs[hex.EncodeToString(vin.Txid)].ID == nil {
			slog.Error("Previous transaction is not correct")
		}
	}
	// 我们需要对修剪部分的交易副本签名，而不是完整的交易
	txCopy := tx.TrimmedCopy() 

	// 遍历交易的输入
	for inID, vin :=range txCopy.VIn {
		// 找到输入引用的（输出所在的）交易
		prevTx := prevTXs[hex.EncodeToString(vin.Txid)]

		// 重置当前输入的签名
		txCopy.VIn[inID].Signature = nil 

		// 输入的公钥（未哈希） 其实就是引用的输出的 公钥哈希值
		txCopy.VIn[inID].PubKey = prevTx.VOut[vin.Vout].PubKeyHash

		dataToSign := []byte(fmt.Sprintf("%x\n", txCopy))
		
		// 利用私钥签名，产生一对数字
		r, s, err := ecdsa.Sign(rand.Reader, &privKey, dataToSign)
		if err != nil {
			slog.Error("failed in signing!")
		}
		// 注意r，s等长，，字节压缩他们生成签名 
		signature := append(r.Bytes(), s.Bytes()...)

		tx.VIn[inID].Signature = signature
		txCopy.VIn[inID].PubKey = nil
	}
}

func(tx *Transaction) Verify(prevTXs map[string]Transaction) bool {
	// TODO: 公钥验证
	// 根据输入引用之前的（输出所在）交易，验证签名是否正确
	if tx.IsCoinbase() {
		// coinbase没有签名，也无需验证
		return true
	}
	for _, vin := range tx.VIn {
		if prevTXs[hex.EncodeToString(vin.Txid)].ID == nil {
			slog.Error("Previous transaction is not correct")
		}
	}
	
	txCopy := tx.TrimmedCopy()
	curve := elliptic.P256()
	fmt.Printf("txCopy: %v\n", txCopy.String())

	for inID, vin := range tx.VIn {
		// 所有的输入都应该要通过验证
		prevTx := prevTXs[hex.EncodeToString(vin.Txid)]
		txCopy.VIn[inID].Signature = nil 
		txCopy.VIn[inID].PubKey = prevTx.VOut[vin.Vout].PubKeyHash
		dataToVerify := []byte(fmt.Sprintf("%x\n", txCopy))
		fmt.Printf("dataToVerify: %x\n", dataToVerify)

		r := big.Int{}
		s := big.Int{}
		sigLen_half := len(vin.Signature) / 2
		r.SetBytes(vin.PubKey[:sigLen_half])
		s.SetBytes(vin.PubKey[sigLen_half:])
		fmt.Printf("%v\n", r.String())
		fmt.Printf("%v\n", s.String())

		x := big.Int{}
		y := big.Int{}
		keyLen_half := len(vin.PubKey) / 2
		x.SetBytes(vin.PubKey[:keyLen_half])
		y.SetBytes(vin.PubKey[keyLen_half:])
		fmt.Printf("%v\n", x.String())
		fmt.Printf("%v\n", y.String())

		rawPubKey := ecdsa.PublicKey{Curve: curve, X: &x, Y: &y}
		fmt.Printf("rawPubKey: %v\n", rawPubKey)

		if !ecdsa.Verify(&rawPubKey, dataToVerify, &r, &s) {
			fmt.Println("Verification fail!")
			return false
		}
		txCopy.VIn[inID].PubKey = nil
	}
	fmt.Println("Verification success!")
	return true
}