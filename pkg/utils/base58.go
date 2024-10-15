package utils

import (
	"bytes"
	"math/big"
)

// Base58 字母表，去掉0OIl
var alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
var base = big.NewInt(58)

func Base58Encode(input []byte) []byte {
	output := []byte{}
	bigInput := BytesToBigInt(input)
	bigZero := big.NewInt(0)
	mod := big.NewInt(0)

	for bigInput.Cmp(bigZero) != 0 {
		bigInput.DivMod(bigInput, base, mod)
		output = append(output, alphabet[mod.Int64()])
	}

	output = reverseBytes(output)
	for b := range input {
		if b == 0x00 {
			output = append([]byte{alphabet[0]}, output...)
		} else {
			break
		}
	}
	return output
}

func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0

	for b := range input {
		if b == 0x00 {
			zeroBytes++
		}
	}

	payload := input[zeroBytes:]
	for _, b := range payload {
		charIndex := bytes.IndexByte(alphabet, b)
		result.Mul(result, base)
		result.Add(result, big.NewInt(int64(charIndex)))
	}

	output := result.Bytes()
	output = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), output...)
	return output
}

func reverseBytes(input []byte) []byte {
	output := []byte{}
	for pos := len(input) - 1; pos >= 0; pos-- {
		output = append(output, input[pos])
	}
	return output
}

func BytesToBigInt(data []byte) *big.Int {
	return big.NewInt(0).SetBytes(data)
}