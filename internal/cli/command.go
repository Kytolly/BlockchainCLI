package cli

import (
	"fmt"
	"os" 
	bm "blockchain/internal/block_model"
	bcm "blockchain/internal/blockchain_model"
	wm "blockchain/internal/wallet_model"
	ts "blockchain/internal/transaction_model"
	st "blockchain/pkg/setting"
	utils "blockchain/pkg/utils"
	"log/slog"
	"strconv"
)

func (cli *CLI) printUsage(){
	fmt.Println("Usage:")
    fmt.Println("blockchain new --address <ADDRESS>  		 				 # To create a blockchain or connect built one, send genesis block to ADDRESS")
    fmt.Println("blockchain createwallet 									 # To generate a new wallet, and save it to wallet file")
	fmt.Println("blockchain getbalance --address <ADDRESS> 				     # To get balance of wallet that located at ADDRESS")
	fmt.Println("blockchain list 											 # To list all addresses in wallet file")
	fmt.Println("blockchain print 										     # To print al blocks in blockchain")
	fmt.Println("blockchain send --from <FROM> --to <TO> --amount <AMOUNT>   # To send <AMOUNT> of bitcoins <FROM> --> <TO>")
	fmt.Println("blockchain log 											 # To display log file content in console")
	fmt.Println("blockchain help 											 # To print usage of CLI commands")
}

func(cli *CLI) validateArgs(){
	// TODO: 验证是否给出命令。
	if len(os.Args) < 2 {
		panic("use command `help` to check out usage")
	}
}

func(cli *CLI) newChain(addressData *string)(*bcm.BlockChain){
	bc := bcm.NewBlockChain(*addressData) 
	defer bc.Close()
	UTXOSet := bcm.UTXOSet{BC: bc}
	UTXOSet.ReIndex()
	return bc
}

func(cli *CLI) getBalance(address string)int{
	if !wm.CheckAddress(address) {
		slog.Warn("Invalid address")
	}
	bc := bcm.GetBlockChain()
	u := bcm.UTXOSet{BC: bc} 
	defer bc.Close()

	balance := 0
	pubKeyHash := utils.GetPubKeyHashInAddress([]byte(address))
	// UTXOs := bc.FindUTXO(pubKeyHash)
	UTXOs := u.FindUTXO(pubKeyHash)
	slog.Debug("UTXOs len", "", len(UTXOs))
	
	for _, out := range UTXOs{
		balance += out.Value
	}
	slog.Info("success", "address", address, "balance", balance)
	fmt.Printf("Balance of %v:%v\n", address, balance)
	return balance
}

func(cli *CLI) createWallet() {
	wallets,_:= wm.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Printf("New address: %s\n", address)
}

func(cli *CLI) listAddresses() {
	wallets,_ := wm.NewWallets()
	addresses := wallets.GetAddress()
	for idx, address := range addresses {
		fmt.Printf("Address %d: %s\n", idx, address)
	}
}

func(cli *CLI) printChain(){
	// TODO：命令行接管打印区块链，利用迭代器遍历整个区块链
	bci := cli.BC.Iterator()
	for {
		block := bci.Next()

		fmt.Printf("============ Block %x ============\n", block.Hash)
		fmt.Printf("Prev. block: %x\n", block.PrevBlockHash)
		pow := bm.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		fmt.Printf("\n\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func(cli *CLI) send(from, to string, amount int) {
	// TODO：命令行接管发送交易，利用区块链的挖矿功能
	if !wm.CheckAddress(from) {
		slog.Warn("Invalid address:", "from", from)
		fmt.Println("the address FROM is invalid, please check it")
	}
	if!wm.CheckAddress(to) {
        slog.Warn("Invalid address:", "to", to)
        fmt.Println("the address TO is invalid, please check it")
    }
	bc := bcm.GetBlockChain() 
	UTXOSet := bcm.UTXOSet{BC: bc}
	defer bc.Close()

	//创建一个通用交易，将挖掘的区块添加到区块链中
	tx := bc.NewUTXOTransaction(from, to, amount, &UTXOSet)
	cbTx := ts.NewCoinbaseTx(from, "")
	txs := []*ts.Transaction{cbTx, tx}
	if tx == nil {
		slog.Info("No transactions to mine")
		slog.Info("Failed Sending!")
		return 
	}

	newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)

	slog.Info("Success sending!")
}

func(cli *CLI) printLog() {
	dest := st.Destination
	err := os.Remove(dest) 
	if err != nil {
		slog.Info(err.Error()) 
	}
}