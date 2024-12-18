package utils

import (
	"bytes"
	"log/slog"
	"math/big"
)

// Base58 字母表，去掉0OIl
var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) []byte {
	//TODO： base58编码
    var result []byte

    x := big.NewInt(0).SetBytes(input)
    base := big.NewInt(int64(len(b58Alphabet)))
    zero := big.NewInt(0)
    mod := &big.Int{}

    for x.Cmp(zero) != 0 {
        x.DivMod(x, base, mod)
        result = append(result, b58Alphabet[mod.Int64()])
    }

    ReverseBytes(result)

    // 修正前导零字节的处理
    for _, b := range input {
        if b == 0x00 {
            result = append([]byte{b58Alphabet[0]}, result...)
        } else {
            break
        }
    }

    return result
}

func Base58Decode(input []byte) []byte {
	// base58解码
    result := big.NewInt(0)
    zeroBytes := 0

    // 修正前导零字节的检查逻辑
    for _, b := range input {
        if b == b58Alphabet[0] {
            zeroBytes++
        } else {
            break
        }
    }

    payload := input[zeroBytes:]
    for _, b := range payload {
        charIndex := bytes.IndexByte(b58Alphabet, b)
        if charIndex < 0 {
            // panic("Invalid Base58 character")
            slog.Error("Invalid Base58 character, maybe check the address!")
        }
        result.Mul(result, big.NewInt(58))
        result.Add(result, big.NewInt(int64(charIndex)))
    }

    decoded := result.Bytes()

    // 处理前导零字节的恢复
    decoded = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decoded...)

    return decoded
}