package server_model

import (
	
)

type getData struct {
	// 对某个区块或交易的请求，只能包含一个区块ID或交易ID
	Addrfrom string // 目标结点的地址
	Type string     // 类型：区块/交易
 	ID []byte       // 区块ID/交易ID，方便在数据库查找
}

func sendGetData(toAddr, t string, id []byte) {
	// TODO： 发送区块哈希/交易哈希
	payload := gobEncode(getData{Addrfrom: nodeAddress, Type: t, ID:id})
    request := append(commandToBytes("getdata"), payload...)
	sendData(toAddr, request)
}



