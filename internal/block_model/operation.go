package block_model

import (
	ts "blockchain/internal/transaction_model"
	wallet "blockchain/internal/wallet_model"
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
)


func(bc *BlockChain) FindUnspentTransactions(address string) []ts.Transaction{
	//TODO：在整个区块链中找到包含能被address锁定的UTXO的交易
	var unspentTXs []ts.Transaction 
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()

	for{
		// 在当前区块查找包含UTXO的交易
		block := bci.Next()

		for _,tx := range block.Transactions{
			// 检查该区块的交易序列中的每一笔交易是否含有UTXO,因为交易存储在区块中
			txID := hex.EncodeToString(tx.ID)

			Outputs:
			for outIdx, out:= range tx.VOut{
				// 准备跳过那些被输入引用的输出
				if spentTXOs[txID] != nil{
					// 被输入引用说明该输出已经被消费
					for _,spentOut := range spentTXOs[txID]{
						if spentOut == outIdx{
							//这个输出不是UTXO
							continue Outputs
						}
					}
				}
				// 确定输入中没有引用这个输出
				if out.CanBeUnlockedWith(address) {
					// 一个output能被address锁定
					unspentTXs = append(unspentTXs, *tx)
				}
			}

			//收集所有可以解锁 被address锁定的输出 的输入
			if !tx.IsCoinbase(){
				// 不是coinbase交易
                for _,in := range tx.VIn{
					// 检查输入是否可以给输出上锁
					if in.CanUnlockOutputWith(address){
						inTxID := hex.EncodeToString(in.Txid)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
					}
				}
			}
		}

		if len(block.PrevBlockHash) == 0{
			break
		}
	}
	return unspentTXs
}

func(bc *BlockChain) FindUTXO(address string) []ts.TXOutput{
	//TODO：为计算余额，还需设计找到能被address锁定的UTXO
	var UTXOs []ts.TXOutput
	UTXs := bc.FindUnspentTransactions(address)
	
	for _, tx := range UTXs{
		for _, out := range tx.VOut{
			if out.CanBeUnlockedWith(address){
				UTXOs = append(UTXOs, out)
			}
		}
	}
	return UTXOs
}

func(bc *BlockChain) FindSpendableOutputs(address string, amount int)(int, map[string][]int){
	//TODO：找到所有未使用过的outputs,准备在新的交易花费它们，这些输出应该都是UTXO
	unspendaleOutputs := make(map[string][]int)
	unspentTXs := bc.FindUnspentTransactions(address)
	acc := 0

	Work:
	for _,tx := range unspentTXs{
		txID :=hex.EncodeToString(tx.ID)
		for outIdx, out:=range tx.VOut{
			if out.CanBeUnlockedWith(address)&& acc < amount{
				acc += out.Value
				unspendaleOutputs[txID] = append(unspendaleOutputs[txID], outIdx)
				if acc >= amount{
					break Work
				}
			}
		}
	}
	return acc, unspendaleOutputs
} 

func(bc *BlockChain) NewUTXOTransaction(from, to string, amount int)*ts.Transaction {
	//TODO:创建一个通用的交易
	var inputs []ts.TXInput
	var outputs []ts.TXOutput

	// 找到所有未使用的outputs确保能够支付给定amount
	acc, validOutputs := bc.FindSpendableOutputs(from, amount)
	if acc < amount {
		slog.Debug("", "acc", acc, "amount", amount)
		slog.Warn("Not enough money to spend!!!")
		fmt.Printf("Not enough money to spend!!!\n")
		return nil
	}

	// 对于每个找到的输出，创建一个引用它的输入
	for txid, outs := range validOutputs{
		txID, _ := hex.DecodeString(txid)

		for _, out := range outs {
			input := ts.TXInput{Txid:txID, Vout:out, ScriptSig:from}
			inputs = append(inputs, input)
		}
	}

	//  每笔交易一般有1-2个输出，如果找到的outputs钱超过了就会找零
	output_send := ts.TXOutput{Value:amount, ScriptPubKtey:to}
	outputs = append(outputs, output_send)
	if acc > amount {
		// 找零钱是因为输出不可分割，超过了就另建一个输出
		output_change := ts.TXOutput{Value:acc-amount, ScriptPubKtey: from}
		outputs = append(outputs, output_change)
	}

	tx := ts.Transaction{ID:nil, VIn:inputs, VOut:outputs}
	tx.SetID()
	// 签名发生在新建UTXO交易中
	bc.SignTransaction(&tx, wallet.PrivateKey)
	return &tx
}

func(bc *BlockChain) FindTransaction(ID []byte)(ts.Transaction, error){
	// TODO：通过ID查找一笔交易
	bci := bc.Iterator()

	for{
		block := bci.Next()

		for _,tx := range block.Transactions{
			if bytes.Compare(tx.ID, ID)==0 {
				// 查找成功，返回
				return *tx, nil
			}
		}

		if len(block.PrevBlockHash)==0{
			break
		}
	}
	return ts.Transaction{}, errors.New("Transaction not found")
}

func(bc *BlockChain) SignTransaction(tx *ts.Transaction, privKey ecdsa.PrivateKey) {
	//TODO: 为交易签名
	// 先找到先前引用输出所在的交易
	prevTXs := bc.FindMapOfPrevTransactions(tx)
	tx.Sign(privKey, prevTXs)
}

func(bc *BlockChain) VerifyTransaction(tx *ts.Transaction)bool{
	prevTXs := bc.FindMapOfPrevTransactions(tx)
	return tx.Verify(prevTXs)
}

func(bc *BlockChain) FindMapOfPrevTransactions(tx *ts.Transaction)map[string]ts.Transaction {
	prevTXs := make(map[string]ts.Transaction)
	for _,vin := range tx.VIn {
		prevTX, err := bc.FindTransaction(vin.Txid)
		if err != nil {
			slog.Info(err.Error())
			continue
		}
		prevTXs[hex.EncodeToString(prevTX.ID)] = prevTX
	}
	return prevTXs
}