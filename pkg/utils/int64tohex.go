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