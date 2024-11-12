package server_model

import (
	bcm "blockchain/internal/blockchain_model"
	bm "blockchain/internal/block_model"
	ts "blockchain/internal/transaction_model"
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"io"
	"log/slog"
	"net"
)

func handleConnection(conn net.Conn, bc *bcm.BlockChain) {
	// TODO：对于网络连接的处理
	request, err := io.ReadAll(conn)
	if err!= nil {
        slog.Error("Failed to read from connection!")
        conn.Close()
        return
    }

	// 读取传过来的具体命令，并进行相应的处理
	command := bytesToCommand(request[:commandLength])
	fmt.Printf("Received %s command\n", command)

	switch command {
	case "version":
		handleVersion(request, bc)
	case "getblocks":
		handleGetBlocks(request, bc)
	case "inv":
		handleInv(request, bc)
    case "getdata":
		handleGetData(request, bc)
	case "block":
		handleBlock(request, bc)
	case "tx":
		handleTx(request, bc)
	default:
		fmt.Println("Unknown command!")
	}
	conn.Close()
}

func handleVersion(request []byte, bc *bcm.BlockChain) {
	// TODO：传输的消息是version命令

	// 按照version结构体解析payload
	var buff bytes.Buffer
	var payload version
	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err!=nil {
		slog.Error(err.Error())
	}
	// 两边进行同步更新
	// 如果本地链短，说明需要获取别人的区块
	// 如果本地链长，说明需要别人需要获取区块，自己把版本发给他让他去别的地方更新区块链
	myBestHeight := bc.GetBestHeight()
	foreignerBestHeight := payload.BestHeight
	if myBestHeight < foreignerBestHeight {
		sendGetBlocks(payload.Addrfrom)
	}else if myBestHeight > foreignerBestHeight {
		sendVersion(payload.Addrfrom, bc)
	}

	// 和别的节点已经通信过，把它加入到knownNodes
	if !nodeIsKnown(payload.Addrfrom) {
		knownNodes = append(knownNodes, payload.Addrfrom)
	}
}

func handleGetBlocks(request []byte, bc *bcm.BlockChain){
	// TODO： 得到所有区块的哈希列表
	var buff bytes.Buffer
	var payload getBlocks
	// 按照getBlocks结构体解析payload
	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err!=nil {
		slog.Error(err.Error())
	}

	// 获取别人的区块实际上先获得哈希列表，减小网络负载
	blocks := bc.GetBlockHashes()
	sendInv(payload.Addrfrom, "block", blocks)
}

func handleInv(request []byte, _ *bcm.BlockChain){
	// TODO： 发送库存的区块哈希列表
	var buff bytes.Buffer
	var payload inventory
	// 按照inventory结构体解析payload
	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err!=nil {
		slog.Error(err.Error())
	}

	// 检查inventory的类型
	// 内存的变量更新可以跟踪那些已经下载的区块，从而实现分布式同步：从不同结点下载区块
	slog.Info(fmt.Sprintf("Received inventory with %d %s\n", len(payload.Items), payload.Type))
	if payload.Type == "block" {
		//要发送的哈希列表保存在内存中的blocksInTransit
		//发送第一个item就删掉它(以及相同的item，其实就是哈希)
		blocksInTransit = payload.Items
		blockHash := payload.Items[0]
		sendGetData(payload.Addrfrom, "block", blockHash)
		// 删除发送过的hash
		newInTransit := [][]byte{}
		for _,b :=range blocksInTransit {
			if !bytes.Equal(b, blockHash) {
				newInTransit = append(newInTransit, b)
			}
		}
		blocksInTransit = newInTransit
	}
	if payload.Type == "tx" {
		// 发送交易
		txID := payload.Items[0]
		txid := hex.EncodeToString(txID)
		if mempool[txid].ID == nil {
			sendGetData(payload.Addrfrom, "tx", txID)
		}
	}
}

func handleGetData(request []byte, bc *bcm.BlockChain){
	// TODO： 处理某个区块或交易的请求，只包含一个区块/交易
	var buff bytes.Buffer
	// 按照getData的方式解析payload
	var payload getData
	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err!=nil {
		slog.Error(err.Error())
	}

	// 获得某个区块
	if payload.Type == "block" {
		block, err := bc.GetBlock([]byte(payload.ID))
		if err != nil {
			slog.Error(err.Error()) 
		}
		sendBlock(payload.Addrfrom, block) 
	}
	// 获得某笔交易
	if payload.Type == "tx" {
		txID := hex.EncodeToString(payload.ID) 
		tx := mempool[txID]
		sendTx(payload.Addrfrom, &tx)
	}
}

func handleBlock(request []byte, bc *bcm.BlockChain){
	//TODO: 处理Block命令，收到真正的区块数据，放入本地的区块链中
	var buff bytes.Buffer
    var payload block
	// 按照block的方式解析payload
    buff.Write(request[commandLength:])
    dec := gob.NewDecoder(&buff)
    err := dec.Decode(&payload)
    if err!= nil {
        slog.Error(err.Error())
    }
	// 反序列化提取出区块，加入区块链
    blockData := payload.Block
	block := bm.DeserializeBlock(blockData)
	fmt.Printf("Received a new block, adding it: %x\n", block.Hash)
	bc.AddBlock(block) // 验证后加入区块链

	
	if len(blocksInTransit) > 0{
		// 如果还有区块没有获取
		// 网络上继续请求发送
		blockHash := blocksInTransit[0]
		sendGetData(payload.Addrfrom, "block", blockHash)
		blocksInTransit = blocksInTransit[1:]
	}else {
		// 收到了全部的区块数据，重新索引UTXOSet
		UTXOSet := bcm.UTXOSet{BC: bc}
		UTXOSet.ReIndex()
	}
}

func handleTx(request []byte, bc *bcm.BlockChain){
	// TODO：处理tx命令，获取真正的交易数据
	var buff bytes.Buffer
	// 按照tx的方式解析payload
    var payload tx
    buff.Write(request[commandLength:])
    dec := gob.NewDecoder(&buff)
    err := dec.Decode(&payload)
    if err!= nil {
        slog.Error(err.Error())
    }

	// 将交易放进内存池中
	txData := payload.Transaction
	tx := ts.DeserializeTransaction(txData)
	txid := hex.EncodeToString(tx.ID)
	mempool[txid] = *tx

	if nodeAddress == knownNodes[0] {
		// 中心结点不会挖矿，向其他结点广播交易
		for _, node := range knownNodes {
			if node != nodeAddress && node != payload.Addrfrom{
				sendInv(node, "tx", [][]byte{tx.ID})
			}
		}
	}else {
		if len(mempool) >=2 && len(miningAddress) > 0 {
		MineTransaction:
		// 矿工结点开始挖矿，打包内存池的交易，生成新的区块，加入区块链
			// 验证交易
			var txs []*ts.Transaction
			for id := range mempool {
				tx := mempool[id]
				if bc.VerifyTransaction(&tx){
					txs = append(txs, &tx)
				}
			}
			if len(txs) == 0 {
				fmt.Println("All transactions are invalid, wait for new transaction...")
				return 
			}
			// 挖矿产生coinbase交易
			cbTx := ts.NewCoinbaseTx(miningAddress, "")
			txs = append(txs, cbTx)

			// 打包成区块
			newBlock := bc.MineBlock(txs)
			UTXOSet := bcm.UTXOSet{BC: bc}
			UTXOSet.ReIndex()
			slog.Info(fmt.Sprintf("New block mined: %x\n", newBlock.Hash))

			// 从内存池中删除打包完成的区块
			for _, tx := range txs {
				txid := hex.EncodeToString(tx.ID)
				delete(mempool, txid)
			}
			// 向周围结点广播这个新区块
			for _, node := range knownNodes {
                if node != nodeAddress {
					sendInv(node, "block", [][]byte{newBlock.Hash})
				}
            }
			if len(mempool)>0 {
				goto MineTransaction
			}
		}
	}
}