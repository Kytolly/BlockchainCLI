package cli

import (
	bm "blockchain/internal/block_model"
	"fmt"
	"log/slog"
)

func(cli *CLI)GetBalance(address string)int{
	cancle, bc := bm.NewBlockChain(address)
	defer func(){
		slog.Info("Shutting down the server...")
		(*cancle)()
	}()

	balance := 0
	UTXOs := bc.FindUTXO(address)
	
	for _, out := range UTXOs{
		balance += out.Value
	}
	slog.Info("success", "address", address, "balance", balance)
	fmt.Printf("Balance of %v:%v\n", address, balance)
	return balance
}