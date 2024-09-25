package setting

import (
	"log/slog"

	"gopkg.in/ini.v1"
)

var initFile = "../../configs/config.ini"

var(
	AppMode 		    string

	LoggerLevel 		string
	Destination 		string

	URI 				string
	DbScheme    		string
	DbHost      		string
	DbPort      		string
	DbTimeOutLimit  	int
	DbName      		string
	DbCollectionName   	string
)

func init(){
	file, err := ini.Load(initFile)
	if err != nil {
		slog.Error("Failed to load config file!")
	}

	loadServer(file)
	loadDatabase(file)
	loadLog(file)
}

func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
}

func loadDatabase(file *ini.File) {
	DbScheme		=file.Section("database").Key("DbScheme").MustString("mongodb")
	DbHost			=file.Section("database").Key("DbHost").MustString("localhost")  
	DbPort			=file.Section("database").Key("DbPort").MustString("27017")  
	DbTimeOutLimit	=file.Section("database").Key("DbTimeOutLimit").MustInt(10)
	DbName			=file.Section("database").Key("DbName").MustString("MyBlockChain")  
	DbCollectionName=file.Section("database").Key("DbCollectionName").MustString("block")

	URI = DbScheme + "://" + DbHost + ":" + DbPort
}

func loadLog(file *ini.File){
	LoggerLevel     = file.Section("log").Key("LoggerLevel").MustString("debug")
    Destination     = file.Section("log").Key("Destination").MustString("console")
}
