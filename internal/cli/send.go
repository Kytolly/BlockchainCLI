package cli

import(
	bm "blockchain/internal/block_model"
	ts "blockchain/internal/transaction_model"
	"log/slog"
)

func(cli *CLI) send(from, to string, amount int) {
	cancle, bc := bm.NewBlockChain(from)
	defer func(){
		slog.Info("Shutting down the server...")
		(*cancle)()
	}()

	//创建一个通用交易，将挖掘的区块添加到区块链中
	//此处不符合比特币的设计规范，还需建立内存池等待矿工挖矿
	tx := bc.NewUTXOTransaction(from, to, amount)
	if tx == nil {
		slog.Info("No transactions to mine")
		slog.Info("Failed Sending!")
		return 
	}
	bc.MineBlock([]*ts.Transaction{tx})
	slog.Info("Success sending!")
}