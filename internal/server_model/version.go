package server_model

import (
	bcm "blockchain/internal/blockchain_model"
	st "blockchain/pkg/setting"
)

const commandLength = st.CommandLength
var nodeVersion = st.Version

type version struct {
	Version 	int     // 版本号
	BestHeight 	int     // 本地区块链的末块的高度
	Addrfrom     string // 本地结点地址
}

func sendVersion(toAddr string, bc *bcm.BlockChain){
	// TODO： 发送的版本信息， 包括区块链末块的高度
	bestHeight := bc.GetBestHeight()
	payload := gobEncode(version{Version: nodeVersion, BestHeight: bestHeight, Addrfrom: nodeAddress})
	request := append(commandToBytes("version"), payload...)
	sendData(toAddr, request)
}