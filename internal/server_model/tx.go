package server_model

import (
	ts "blockchain/internal/transaction_model"
)

type tx struct {
	Addrfrom string
	Transaction []byte
}

func sendTx(addrfrom string, tx *ts.Transaction) {
	// TODO： 发送真正的交易数据 
	payload := gobEncode(block{Addrfrom: addrfrom, Block: tx.Serialize()})
    request := append(commandToBytes("block"), payload...)
	sendData(addrfrom, request)
}
func SendTx_center(tx *ts.Transaction){
	sendTx(knownNodes[0], tx)
}