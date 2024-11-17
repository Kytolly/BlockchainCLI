package server_model

import(
	bm "blockchain/internal/block_model"
)

type block struct {
	Addrfrom string
	Block []byte
}

func sendBlock(toAddr string, bkd *bm.Block) {
	// TODO： 发送区块数据
	payload := gobEncode(block{Addrfrom: nodeAddress, Block: bkd.Serialize()})
    request := append(commandToBytes("block"), payload...)
	sendData(toAddr, request)
}