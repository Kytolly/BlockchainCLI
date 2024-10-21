package transaction_model

import (
	"bytes"
	"encoding/gob"
	"log/slog"
)

type TXOutputs struct {
	Outputs []TXOutput 
}

func(os TXOutputs) Serialize() []byte {
	// TODO： 序列化交易输出序列。
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(os)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

func DeserializeOutputs(data []byte) *TXOutputs {
	// TODO: 反序列化交易输出序列
	var txo TXOutputs

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&txo)
	if err != nil{
		slog.Error(err.Error())
	}
	return &txo
}