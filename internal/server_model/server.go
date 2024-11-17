package server_model

import (
	bcm "blockchain/internal/blockchain_model"
	st "blockchain/pkg/setting"
	"fmt" 
	"log/slog"
	"net"
)

var nodeAddress string  // 本地地址
var miningAddress string  
var knownNodes = []string {st.CenterPort}
var protocol = st.Protocol

var blocksInTransit [][]byte

func StartServer(nodeID, minerAddress string) {
	// TODO: 接受消息的服务器
	fmt.Println("Preparing to start server...")

	nodeAddress = fmt.Sprintf("localhost:%s", nodeID)
	miningAddress = minerAddress
	fmt.Printf("the local node address is %s\n", nodeAddress)

	// 建立本地服务器
	ln, err := net.Listen(protocol, nodeAddress)
	if err!= nil {
		slog.Error("Failed to listen starting server !")
		return
	}
	defer ln.Close()
	fmt.Println("Local Server Listening...")

	// 本地新建区块链，有则退出，向knownNodes发送版本信息，同步更新
	bc := bcm.GetBlockChain(nodeID)
	if nodeAddress != knownNodes[0]{
		sendVersion(knownNodes[0], bc)
	}
	fmt.Println("Local Blockchain ready...")

	fmt.Println("Accepting message...")
	for {
		connect, err := ln.Accept()
		if err != nil {
			slog.Error("Failed to accept !")
		}
		go handleConnection(connect, bc)
	}
}

