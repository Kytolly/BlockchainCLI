package server_model

import(
	// bcm "blockchain/internal/blockchain_model"
)

type getBlocks struct {
	// 命令getBlocks的含义是获得你的区块哈希列表
	Addrfrom string  // 发送的目标结点
}

func sendGetBlocks(addrfrom string){
	// TODO： 发送区块哈希列表
	payload := gobEncode(getBlocks{Addrfrom: addrfrom})
    request := append(commandToBytes("getblocks"), payload...)
	sendData(addrfrom, request)
}