package utils

import(
	"bytes"
	"fmt"
	"testing"
)

func TestBase58(t *testing.T) {
	cases := []byte("hello_world")
	encoded := Base58Encode(cases)
	decoded := Base58Decode(encoded)
	if !bytes.Equal(cases, decoded) {
		fmt.Printf("cases: %x\n", cases)
		fmt.Printf("encoded: %x\n", encoded)
		fmt.Printf("decoded: %x\n", decoded)
		panic("Base58 encoding/decoding failed")
	}
}
