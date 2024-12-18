package setting

import (
	"log/slog"
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var initFile = "/home/kytolly/Project/gowork/blockchain/configs/config.ini"

const CommandLength int = 12

var(
	AppMode 		    string
	Protocol  			string
	CenterPort 			string

	LoggerLevel 		string
	Destination 		string

	URI 				string
	DbScheme    		string
	DbHost      		string
	DbPort      		string
	DbTimeOutLimit  	int
	DbName      		string
	DbCollectionName   	string

	Version  			int
	ChecksumLen  		int
	WalletFile          string

	Subsidy             int

	TargetBits          int

	GenesisCoinbaseData string
	ChainDbFile 		string	
	BlockBucket         string
	UTXOBucketName      string

	NODE_ID             string
)

func init(){
	file, err := ini.Load(initFile)
	if err != nil {
		slog.Error("Failed to load config file!")
	}

	loadEnv()
	loadServer(file)
	// loadDatabase(file)
	loadWallet(file)
	loadTransaction(file)
	loadLog(file)
	loadPow(file)
	loadBlockchain(file)
}

func loadEnv(){
	NODE_ID = os.Getenv("NODE_ID")
	if NODE_ID == "" {
		fmt.Println("NODE_ID env. var is not set!")
		os.Exit(1)
	}
}

func loadServer(file *ini.File) {
	AppMode 	= file.Section("server").Key("AppMode").MustString("debug")
	Protocol 	= file.Section("server").Key("Protocol").MustString("tcp")	
	CenterPort 	= file.Section("server").Key("CenterPort").MustString("localhost:3000")
	// CommandLength = file.Section("server").Key("CommandLength").MustInt(12)
}

// func loadDatabase(file *ini.File) {
// 	DbScheme		=file.Section("database").Key("DbScheme").MustString("mongodb")
// 	DbHost			=file.Section("database").Key("DbHost").MustString("localhost")  
// 	DbPort			=file.Section("database").Key("DbPort").MustString("27017")  
// 	DbTimeOutLimit	=file.Section("database").Key("DbTimeOutLimit").MustInt(10)
// 	DbName			=file.Section("database").Key("DbName").MustString("MyBlockChain")  
// 	DbCollectionName=file.Section("database").Key("DbCollectionName").MustString("block")

// 	URI = DbScheme + "://" + DbHost + ":" + DbPort
// }

func loadWallet(file *ini.File) {
	Version 		 = file.Section("wallet").Key("Version").MustInt(1)
	ChecksumLen 	 = file.Section("wallet").Key("ChecksumLen").MustInt(10)
	WalletFile       = file.Section("wallet").Key("WalletFile").MustString("wallet.dat")

	WalletFile = fmt.Sprintf(WalletFile, NODE_ID)
}

func loadTransaction(file *ini.File) {
	Subsidy = file.Section("transaction").Key("Subsidy").MustInt(50)
}

func loadLog(file *ini.File){
	LoggerLevel     = file.Section("log").Key("LoggerLevel").MustString("debug")
    Destination     = file.Section("log").Key("Destination").MustString("console")
}

func loadPow(file *ini.File){
	TargetBits = file.Section("pow").Key("TargetBits").MustInt(5)
}

func loadBlockchain(file *ini.File){
	GenesisCoinbaseData = file.Section("blockchain").Key("GenesisCoinbaseData").MustString("Genesis Block!")
	ChainDbFile         = file.Section("blockchain").Key("ChainDbFilename").MustString("chain_%s.db")
	BlockBucket         = file.Section("blockchain").Key("BlockBucket").MustString("block")
	UTXOBucketName 	 	= file.Section("blockchain").Key("UTXOBucket").MustString("chainstate")

	// ChainDbFile = fmt.Sprintf(ChainDbFile, NODE_ID)
}