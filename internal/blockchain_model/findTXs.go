package blockchain_model

import (
	ts "blockchain/internal/transaction_model"
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
)

func(bc *BlockChain) FindUnspentTransactions(pubKeyHash []byte) []ts.Transaction{
	//TODO：在整个区块链中找到包含能被公钥锁定的UTXO的交易
	var unspentTXs []ts.Transaction 
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()
	slog.Debug("", "now PubKeyHash", fmt.Sprintf("%x", pubKeyHash))

	for{
		// 在当前区块查找包含UTXO的交易
		block := bci.Next()
		slog.Debug("block.transactions", "len", len(block.Transactions))
		for _,tx := range block.Transactions{
			// 检查该区块的交易序列中的每一笔交易是否含有UTXO,因为交易存储在区块中
			txID := hex.EncodeToString(tx.ID)
			slog.Debug("tx in block.Transaction", "txID_string", txID)

			Outputs:
			for outIdx, out:= range tx.VOut{
				// 准备跳过那些被输入引用的输出
				slog.Debug("in Outputs lable", "outIdx", outIdx)
				if spentTXOs[txID] != nil{
					// 被输入引用说明该输出已经被消费
					for _,spentOut := range spentTXOs[txID]{
						slog.Debug("in Outputs lable", "spentOut", spentOut)
						if spentOut == outIdx{
							//这个输出不是UTXO
							slog.Debug("continue Outputs lable", "spentOut == outIdx", outIdx)
							continue Outputs
						}
					}
				}
				// 确定输入中没有引用这个输出
				slog.Debug("check out", "out PubKeyHash", fmt.Sprintf("%x", out.PubKeyHash))
				if out.IsLockedWithKey(pubKeyHash) {
					// 一个output能被公钥锁定
					slog.Debug("out islockedwithkey", "pubKeyHash", pubKeyHash)
					unspentTXs = append(unspentTXs, *tx)
				}
			}

			//收集所有可以解锁 被address锁定的输出 的输入
			if !tx.IsCoinbase(){
				// 不是coinbase交易
				slog.Debug("tx is not coinbase")
                for _,in := range tx.VIn{
					// 检查输入是否可以给输出解锁
					slog.Debug("in PublicKey")
					if in.UsesKey(pubKeyHash){
						slog.Debug("in UsesKey pubKeyHash", "pubKeyHash", pubKeyHash)
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
	slog.Debug("length of findUnspent transactions", "len", len(unspentTXs))
	return unspentTXs
}

// func(bc *BlockChain) FindUTXO(pubKeyHash []byte) []ts.TXOutput{
// 	//TODO：为计算余额，还需设计找到能被address锁定的UTXO
// 	var UTXOs []ts.TXOutput
// 	UTXs := bc.FindUnspentTransactions(pubKeyHash)
	
// 	for _, tx := range UTXs{
// 		for _, out := range tx.VOut{
// 			if out.IsLockedWithKey(pubKeyHash) {
// 				UTXOs = append(UTXOs, out)
// 			}
// 		}
// 	}
// 	return UTXOs
// }

// func(bc *BlockChain) FindSpendableOutputs(pubKeyHash []byte, amount int)(int, map[string][]int){
// 	//TODO：找到所有未使用过的outputs,准备在新的交易花费它们，这些输出应该都是UTXO
// 	unspendaleOutputs := make(map[string][]int)
// 	unspentTXs := bc.FindUnspentTransactions(pubKeyHash)
// 	acc := 0

// 	Work:
// 	for _,tx := range unspentTXs{
// 		txID :=hex.EncodeToString(tx.ID)
// 		for outIdx, out:=range tx.VOut{
// 			if out.IsLockedWithKey(pubKeyHash)&& acc < amount{
// 				acc += out.Value
// 				unspendaleOutputs[txID] = append(unspendaleOutputs[txID], outIdx)
// 				if acc >= amount{
// 					break Work
// 				}
// 			}
// 		}
// 	}
// 	return acc, unspendaleOutputs
// } 

func(bc *BlockChain) FindTransaction(ID []byte)(ts.Transaction, error){
	// TODO：通过ID查找一笔交易
	bci := bc.Iterator()

	for{
		block := bci.Next()

		for _,tx := range block.Transactions{
			if bytes.Equal(tx.ID, ID) {
				// 查找成功，返回
				return *tx, nil
			}
		}

		if len(block.PrevBlockHash)==0{
			break
		}
	}
	return ts.Transaction{}, errors.New("transaction not found")
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

func(bc *BlockChain) FindUTXO() map[string]ts.TXOutputs {
	return nil
}