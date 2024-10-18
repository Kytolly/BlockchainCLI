package utils

import (
	"bytes"
	"encoding/binary"
)

func Int64ToHex(x int64) []byte{
	buff := new(bytes.Buffer)
	err  := binary.Write(buff, binary.BigEndian, x)
	if err != nil {
        panic(err)
		// log.Fatal(err)
    }
	return buff.Bytes()
} 

func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}