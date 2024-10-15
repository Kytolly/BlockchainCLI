package block_model

import (
	"blockchain/pkg/utils"
	"bytes"
	"crypto/sha256"
	"fmt" 
	"log/slog"
	"math"
	"math/big"
)

// 定义挖矿的难度
const targetBits = 5
// 定义 nonce 的上限
const maxNonce = math.MaxInt64

type ProofOfWork struct{
	block 	*Block
	target 	*big.Int // 一个内存占用少于 256 位的目标
}

func NewProofOfWork(b *Block) *ProofOfWork {
	// TODO: 创建一个新的工作量证明
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits)) // 左移 targetBits 位

	return &ProofOfWork{b, target}
}

func(pow *ProofOfWork) prepareData(nonce int) []byte{
	// TODO: 组装需要进行 SHA256 哈希的数据作为hashcash的计数器
	// 工作量证明必须考虑存储在区块中的交易，保证区块链作为交易存储的一致性和可靠性
	data := bytes.Join(
		[] []byte{
			pow.block.PrevBlockHash,
            pow.block.HashTransactions(),
            utils.Int64ToHex(pow.block.Timestamp),
            utils.Int64ToHex(int64(nonce)),
			utils.Int64ToHex(int64(targetBits)),
		},
		[]byte{},
	)
	return data
}

func(pow *ProofOfWork) Run() (int, []byte) {
	// TODO: 进行工作量证明，返回找到的 nonce 和 hash
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	// fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	// slog.Info("Mining the block containing: ","data",  pow.block.Data)

	for nonce < maxNonce {
		data := pow.prepareData(nonce)
        hash = sha256.Sum256(data)

        hashInt.SetBytes(hash[:])

        if hashInt.Cmp(pow.target) == -1 {
			// hashInt < target，证明找到了符合条件的nonce
            break
        } else {
            nonce++
        }

		slog.Info("(Proof of work)", "Hash", fmt.Sprintf("%x", hash), "NONCE", nonce)
	}
	slog.Info("A Block is mined successfully!")
	fmt.Printf("A Block is mined successfully!\n")

	return nonce, hash[:]
}


func(pow *ProofOfWork) Validate() bool{
	// TODO: 对结果进行验证，看是否满足工作量证明难度。
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := (hashInt.Cmp(pow.target) == -1)
	return isValid
}