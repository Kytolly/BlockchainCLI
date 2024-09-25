package main

import (
	"blockchain/internal/cli"
	"blockchain/pkg/mid"
	bm "blockchain/internal/block_model"
	"log/slog"
	//"context"
)
func main() {
	slog.SetDefault(mid.L)
	
	cancle, bc :=bm.NewBlockChain()
	defer func(){
		slog.Info("Shutting down the server...")
		(*cancle)()
	}()

	cli := cli.CLI{BC: bc}
	cli.Run()
}