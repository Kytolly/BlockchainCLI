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
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block Data")

	// 创建两个子命令
	switch os.Args[1]{
	case "addblock":
		err = addBlockCmd.Parse(os.Args[2:])
	case "printchain":
		err = printChainCmd.Parse(os.Args[2:])
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
        cli.printChain()
    }
}