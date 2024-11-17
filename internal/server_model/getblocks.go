package server_model

import(
	// bcm "blockchain/internal/blockchain_model"
)

type getBlocks struct {
	// 命令getBlocks的含义是获得你的区块哈希列表
	Addrfrom string  // 告诉对方自己的地址
}

func sendGetBlocks(toAddr string){
	// TODO： 发送希望获得区块哈希列表的命令
	payload := gobEncode(getBlocks{Addrfrom: nodeAddress})
    request := append(commandToBytes("getblocks"), payload...)
	sendData(toAddr, request)
}