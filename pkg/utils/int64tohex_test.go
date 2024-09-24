package utils

import (
	"bytes"
	"fmt"
	"testing"
)

func TestInt64ToHex(t *testing.T) {
    input := int64(1891915618515156485)
    expected := []byte{0x1a, 0x41, 0x6f, 0x3b, 0x32, 0x9b, 0xfe, 0x05}
	fmt.Printf("Input: %v\n", input)
	fmt.Printf("Expected: %v\n", expected)

	results := Int64ToHex(input)
	fmt.Printf("Results: %v\n", results)

	if !bytes.Equal(results, expected) {
        t.Fatalf("Test failed: Expected %v, got %v", expected, results)
    }

    fmt.Println("Test passed")
}