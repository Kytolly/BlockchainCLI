package transaction_model

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