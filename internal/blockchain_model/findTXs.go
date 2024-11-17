package blockchain_model

import (
	ts "blockchain/internal/transaction_model"
	"bytes"
	"encoding/hex"
	"errors"
	"log/slog"
)

func(bc *BlockChain) FindUTXO() map[string]ts.TXOutputs{ 
	//TODO：在整个区块链中找到未花费的交易输出 
	spentTXOs := make(map[string][]int) 
	UTXO := make(map[string]ts.TXOutputs)
	bci := bc.Iterator() 

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
				outs := UTXO[txID]
				outs.Outputs = append(outs.Outputs, out)
				UTXO[txID] = outs
				slog.Debug("outs", "len", len(outs.Outputs))
			}

			//收集所有可以解锁 被address锁定的输出 的输入
			if !tx.IsCoinbase(){
				// 不是coinbase交易
				slog.Debug("tx is not coinbase")
                for _,in := range tx.VIn{ 
					inTxID := hex.EncodeToString(in.Txid)
					spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
				}
			}
		}

		if len(block.PrevBlockHash) == 0{
			break
		}
	}
	slog.Debug("length of UTXO", "len", len(UTXO))
	return UTXO
}

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
		}
		prevTXs[hex.EncodeToString(prevTX.ID)] = prevTX
	}
	slog.Debug("find map of previous transactions", "len", len(prevTXs))
	return prevTXs
}