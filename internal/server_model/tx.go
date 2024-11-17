package server_model

import (
	ts "blockchain/internal/transaction_model"
	st "blockchain/pkg/setting"
	"fmt"
)

type tx struct {
	Addrfrom string
	Transaction []byte
}

func sendTx(toAddr string, txd *ts.Transaction) {
	// TODO： 发送真正的交易数据 
	payload := gobEncode(tx{Addrfrom: nodeAddress, Transaction: txd.Serialize()})
    request := append(commandToBytes("tx"), payload...)
	sendData(toAddr, request)
}
func SendTx_center(txd *ts.Transaction){
	// TDDO:给中心结点发送交易，让中心结点广播
	id := st.NODE_ID
	node := fmt.Sprintf("localhost:%s", id)
	payload := gobEncode(tx{Addrfrom: node, Transaction: txd.Serialize()})
    request := append(commandToBytes("tx"), payload...)
	fmt.Printf("send tx(%x) from %s to %s\n", txd, node, knownNodes[0])
	sendData(knownNodes[0], request)
}