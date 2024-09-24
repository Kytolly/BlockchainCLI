package main

import (
	"blockchain/internal/cli"
	bm "blockchain/internal/block_model"
	"log/slog"
	//"context"
)
func main() {
	cancle, bc :=bm.NewBlockChain()
	defer func(){
		slog.Info("Shutting down the server...")
		(*cancle)()
	}()

	cli := cli.CLI{BC: bc}
	cli.Run()
}