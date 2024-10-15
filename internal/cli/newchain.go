package cli

import(
	"context"
	bm "blockchain/internal/block_model"
	"log/slog"
)

func(cli *CLI) NewChain(addressData *string)(*context.CancelFunc, *bm.BlockChain){
	cancle, bc := bm.NewBlockChain(*addressData)
	defer func(){
		slog.Info("Shutting down the server...")
		(*cancle)()
	}()
	return cancle, bc
}