package server_model

type inventory struct { 
	// 实现广播机制，将区块/交易的哈希列表广播给其他节点
	Addrfrom string // 发送的目标节点
	Type string    // 表示区块还是交易
	Items [][]byte // 区块/交易的哈希列表,这就是库存的含义
}

func sendInv(addrfrom, t string, items [][]byte) {
    // TODO: 发送inv
	payload := gobEncode(inventory{Addrfrom: addrfrom, Type: t, Items: items})
    request := append(commandToBytes("inv"), payload...)
	sendData(addrfrom, request)
}

