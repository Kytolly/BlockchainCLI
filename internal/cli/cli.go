package cli

import (
	// bm "blockchain/internal/block_model"
	bcm "blockchain/internal/blockchain_model"
	st "blockchain/pkg/setting"
	"flag"
	"fmt"
	"log/slog"
	"os"
)

type CLI struct{
	BC *bcm.BlockChain
}
var nodeID = st.NODE_ID

func(cli *CLI) Run(){
	fmt.Println("Welcome to the blockchain CLI!")

	var err error
	cli.validateArgs()

	// 解析命令行参数
	newCmd:= flag.NewFlagSet("new", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createwalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	printCmd := flag.NewFlagSet("print", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	reindexCmd := flag.NewFlagSet("reindex", flag.ExitOnError)
	startCmd := flag.NewFlagSet("start", flag.ExitOnError)
	logCmd := flag.NewFlagSet("log", flag.ExitOnError)
	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)

	addressData := newCmd.String("address", "Defalut Address", "The Address creating chain")
	addressBalanceData := getBalanceCmd.String("address", "Defalut Address", "The Address creating chain")
	fromData := sendCmd.String("from", "Defalut From Address", "The Address from")
	amountData := sendCmd.Int("amount", 0, "the Amount of sendCmd")
	toData := sendCmd.String("to", "Defalut To Address", "The Address to") 
	mineData := sendCmd.Bool("mine", false, "Mine immediately on the same node")
	startData := startCmd.String("miner", "", "Enable mining mode and send reward to ADDRESS")
	logRemoveData := logCmd.Bool("remove", false, "Determined to remove log?")

	// 创建子命令
	switch os.Args[1]{
	case "new":
		err = newCmd.Parse(os.Args[2:])
	case "getbalance":
		err = getBalanceCmd.Parse(os.Args[2:])
	case "createwallet":
		err = createwalletCmd.Parse(os.Args[2:])
    case "list":
        err = listCmd.Parse(os.Args[2:])
	case "print":
		err = printCmd.Parse(os.Args[2:])
	case "reindex":
		err = reindexCmd.Parse(os.Args[2:])
    case "start":
		err = startCmd.Parse(os.Args[2:])
	case "send":
		err = sendCmd.Parse(os.Args[2:])
	case "log":
		err = logCmd.Parse(os.Args[2:])
    case "help":
        err = helpCmd.Parse(os.Args[2:])
	default:
		cli.printUsage()
		os.Exit(1)
	}
	if err != nil {
		slog.Error(err.Error())
	}


	if newCmd.Parsed() {
		if *addressData == ""{
			newCmd.Usage()
            os.Exit(1)
		}
		bc:= cli.newChain(addressData)
		cli.BC = bc
	}

	if getBalanceCmd.Parsed(){
		if *addressData == ""{
			getBalanceCmd.Usage()
            os.Exit(1)
		}
		cli.getBalance(*addressBalanceData)
	}

	if createwalletCmd.Parsed() {
		cli.createWallet(nodeID)
	}

	if listCmd.Parsed(){
		cli.listAddresses(nodeID)
	}

	if printCmd.Parsed(){
		cli.printChain(nodeID)
	}

	if reindexCmd.Parsed() {
		cli.reIndexUTXO(nodeID)
	}

	if sendCmd.Parsed() {
		if *fromData == "" || *toData == "" || *amountData<=0 {
            sendCmd.Usage()
            os.Exit(1)
        }
        cli.send(*fromData, *toData, *amountData, nodeID, *mineData) 
	}

	if startCmd.Parsed() {
		nodeID := os.Getenv("NODE_ID")
		if nodeID == "" {
			startCmd.Usage()
			os.Exit(1)
		}
		cli.start(nodeID, *startData)
	}
    if logCmd.Parsed(){
		if *logRemoveData {
			dest := st.Destination
            os.Remove(dest)
        } else {
            cli.printLog()
        }
    }

	if helpCmd.Parsed(){
        cli.printUsage()
    }
}