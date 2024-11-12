package transaction_model

import (
	"bytes"
	"encoding/gob"
	"log/slog"
)

func(tx *Transaction) TrimmedCopy() Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	for _, vin := range tx.VIn {
		inputs = append(inputs, TXInput{Txid: vin.Txid, Vout: vin.Vout, Signature: nil, PubKey: nil})
	}

	for _, vout := range tx.VOut {
		outputs = append(outputs, TXOutput{Value: vout.Value, PubKeyHash: vout.PubKeyHash})
	}

	txCopy := Transaction{
		ID: tx.ID, 
		VIn: inputs, 
		VOut: outputs,
	}
	return txCopy
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

func DeserializeTransaction(data []byte) (*Transaction) {
	// TODO: 反序列化交易
	var tx Transaction
    decoder := gob.NewDecoder(bytes.NewReader(data))
    err := decoder.Decode(&tx)
    if err!= nil {
		slog.Error("Error decoding transaction")
        return nil
    }

    return &tx 
}