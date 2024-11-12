package server_model

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log/slog"
	"net"
)

func commandToBytes(command string) []byte {
	// TODO: 创建一个 12 字节的缓冲区，并用命令名称填充它，将其余字节留空
    var bytes [commandLength]byte
	for i,c := range command {
		bytes[i] = byte(c)
	}
	return bytes[:]
}

func bytesToCommand(bytes []byte) string {
	// TODO: 提取命令名称并使用正确的处理程序处理
	var command []byte
	for _, b := range bytes {
		if b != 0x00 {
			command = append(command, b)
		}
	}
	return fmt.Sprintf("%x", command)
}


func nodeIsKnown(addrfrom string) bool {
	for _, node := range knownNodes {
		if node == addrfrom {
			return true
		}
	}
	return false
}

func gobEncode(data interface{}) []byte {
	// TODO: 序列化某个对象(version,getblocks)，准备payload
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(data)
	if err != nil{
		slog.Info(err.Error())
	}
	return result.Bytes()
}

func sendData(addr string, data []byte) {
	conn, err := net.Dial(protocol, addr)
	if err != nil {
		slog.Info(addr+"is not available now!")
		var updatedNodes []string

		// 在knownNodes中删掉addr,因为通信即将结束
		for _,node := range knownNodes {
			if node != addr{
				updatedNodes = append(updatedNodes, node)
			}
 		}
		knownNodes = updatedNodes
		return 
	}
	defer conn.Close()

	_, err = io.Copy(conn, bytes.NewReader(data))
	if err != nil{
		slog.Error(err.Error())
	}
}