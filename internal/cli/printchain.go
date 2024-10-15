package cli

import (
	"fmt"
	bm "blockchain/internal/block_model"
	"strconv"
)

func (cli *CLI) printChain(){
	// TODO：命令行接管打印区块链，利用迭代器遍历整个区块链
	bci := cli.BC.Iterator()
	for {
		block := bci.Next()

		fmt.Printf("Prev.hash: %x\n", block.PrevBlockHash)
		//fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := bm.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}