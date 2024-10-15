package cli

import (
	bm "blockchain/internal/block_model"
	"flag"
	"fmt"
	"log/slog"
	"os"
)

type CLI struct{
	BC *bm.BlockChain
}

func(cli *CLI) Run(){
	slog.Info("Welcome to the blockchain CLI! Let's start with help.")
	// cli.printUsage()

	var err error
	cli.validateArgs()

	// 解析命令行参数
	createChainCmd:= flag.NewFlagSet("new", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("new", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	addressData := createChainCmd.String("address", "Defalut Address", "The Address creating chain")
	addressBalanceData := getBalanceCmd.String("address", "Defalut Address", "The Address creating chain")
	fromData := sendCmd.String("from", "Defalut From Address", "The Address from")
	amountData := sendCmd.Int("amount", 0, "the Amount of sendCmd")
	toData := sendCmd.String("to", "Defalut To Address", "The Address to")
	//addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	//addBlockData := addBlockCmd.String("data", "", "Block Data")

	// 创建子命令
	switch os.Args[1]{
	case "new":
		err = createChainCmd.Parse(os.Args[2:])
	case "getbalance":
		err = getBalanceCmd.Parse(os.Args[2:])
	case "send":
		err = sendCmd.Parse(os.Args[2:])
	case "printchain":
		err = printChainCmd.Parse(os.Args[2:])
	// case "addblock":
	// 	err = addBlockCmd.Parse(os.Args[2:])
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	if createChainCmd.Parsed() {
		if *addressData == ""{
			createChainCmd.Usage()
            os.Exit(1)
		}
		_, bc:= cli.NewChain(addressData)
		cli.BC = bc
	}

	if sendCmd.Parsed() {
		if *fromData == "" || *toData == ""{
            sendCmd.Usage()
            os.Exit(1)
        }
        cli.send(*fromData, *toData, *amountData) 
	}

	if getBalanceCmd.Parsed(){
		if *addressData == ""{
			createChainCmd.Usage()
            os.Exit(1)
		}
		cli.GetBalance(*addressBalanceData)
	}

	if printChainCmd.Parsed() {
        cli.printChain()
    }

	// 旧参数，不再解析
	// if addBlockCmd.Parsed() {
	// 	if *addBlockData == "" {
	// 		addBlockCmd.Usage()
	// 		os.Exit(1)
	// 	}
	// 	cli.addBlock(*addBlockData)
	// }
}