package blockchain_model

import (
	st "blockchain/pkg/setting"
	ts "blockchain/internal/transaction_model"
	bm "blockchain/internal/block_model"
	"encoding/hex"
	"log/slog"

	"github.com/boltdb/bolt"
)

type UTXOSet struct {
	BC *BlockChain
}

var utxoBucket = st.UTXOBucketName

func(u UTXOSet) ReIndex() {
	// 缓存UTXO集合
	db := u.BC.db
	bucketName := []byte(utxoBucket)

	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket(bucketName)
		if err != nil {
			slog.Info(err.Error())
		}
		_, err = tx.CreateBucket(bucketName)
		if err != nil {
			slog.Info(err.Error())
		}
		return err
	})
	if err != nil {
		slog.Error(err.Error())
        return
	}
	UTXO := u.BC.FindUTXO() 

	err = db.Update(func(tx *bolt.Tx)error {
		b := tx.Bucket(bucketName)
		var err error
		for txID, outs := range UTXO {
			key, err := hex.DecodeString(txID)
			if err != nil {
				slog.Error(err.Error())
                continue
			}
			err = b.Put(key, outs.Serialize())
			if err != nil {
				slog.Error(err.Error()) 
			}
		}
		return err 
	})
	if err != nil {
		slog.Error(err.Error()) 
	}
}

func(u UTXOSet) FindSpendableOutputs(pubKeyHash []byte, amount int)(int, map[string][]int){
	//TODO：找到所有未使用过的outputs,准备在新的交易花费它们，这些输出应该都是UTXO
	unspendaleOutputs := make(map[string][]int) 
	acc := 0
	db := u.BC.db

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxoBucket))
		c := b.Cursor()

		for k, v := c.First(); k!=nil; k,v = c.Next() {
			txID := hex.EncodeToString(k)
			outs := ts.DeserializeOutputs(v)

			for outIdx, out := range outs.Outputs {
				if out.IsLockedWithKey(pubKeyHash) && acc < amount {
					acc += out.Value
					unspendaleOutputs[txID] = append(unspendaleOutputs[txID], outIdx)
				}
			}
		}
		return nil
	})
	if err!= nil {
        slog.Error(err.Error())
    }

	return acc, unspendaleOutputs
} 

func(u UTXOSet) FindUTXO(pubKeyHash []byte) []ts.TXOutput {
	var UTXOs []ts.TXOutput 
	db := u.BC.db

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxoBucket))
		c := b.Cursor()

		for k, v:= c.First(); k!=nil; k,v = c.Next() {
			outs := ts.DeserializeOutputs(v)
			slog.Debug("now outs", "len", len(outs.Outputs))
			
			for _,out := range outs.Outputs {
				if out.IsLockedWithKey(pubKeyHash){
					UTXOs = append(UTXOs, out)
				}
			}
		}
		return nil
	})
	if err != nil {
		slog.Error(err.Error())
	}

	return UTXOs
}

func(u UTXOSet) Update(block *bm.Block) {
	// TODO： 开采新区块时，应该更新UTXO集
	// 在新开采的交易中，删除花费的输出，添加未使用的输出
	db := u.BC.db

	err := db.Update(func(tx *bolt.Tx)error {
		b := tx.Bucket([]byte(utxoBucket))
		var err error

		for _, tx := range block.Transactions {
			if !tx.IsCoinbase() {
				for _, vin := range tx.VIn {
					updateOuts := ts.TXOutputs {}
					outBytes := b.Get(vin.Txid)
					outs := ts.DeserializeOutputs(outBytes)

					for outIdx, out := range outs.Outputs {
						if outIdx != vin.Vout {
							updateOuts.Outputs = append(updateOuts.Outputs, out)
						}
					}

					if len(updateOuts.Outputs) == 0 {
						err = b.Delete(vin.Txid)
					}else {
						err = b.Put(vin.Txid, updateOuts.Serialize())
					}
					if err != nil {
						slog.Error(err.Error())
					}
				}
			}
			newOutputs := ts.TXOutputs{}
			newOutputs.Outputs = append(newOutputs.Outputs, tx.VOut...)
			err = b.Put(tx.ID, newOutputs.Serialize())
		}
		return err
	})
	if err != nil {
		slog.Error(err.Error())
	}
}