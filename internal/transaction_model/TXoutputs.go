package transaction_model

import (

)

type TXOutputs struct {
	Outputs []TXOutput 
}

func(os TXOutputs) Serialize() []byte {
	return []byte{}
}

func DeserializeOutputs(outBytes []byte) *TXOutputs {
	return &TXOutputs{}
}