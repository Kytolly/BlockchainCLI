package server_model

import(
	bm "blockchain/internal/block_model"
)

type block struct {
	Addrfrom string
	Block []byte
}

func sendBlock(addrfrom string, bk *bm.Block) {
	// TODO： 发送区块数据
	payload := gobEncode(block{Addrfrom: addrfrom, Block: bk.Serialize()})
    request := append(commandToBytes("block"), payload...)
	sendData(addrfrom, request)
}