package mid

import (
	"log/slog"
	"os"
	"fmt" 
	st"blockchain/pkg/setting"
)

var L *slog.Logger
var mapLevelStr = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func init() {
	Level, ok:= mapLevelStr[st.LoggerLevel]
	if !ok {
		fmt.Printf("unknown logger level: %s", st.LoggerLevel)
		return 
	}
	Options := &slog.HandlerOptions{
		AddSource:  	true, // 记录日志位置
		Level: 			Level, // 设置日志等级
		ReplaceAttr: 	nil,
	}

	switch st.Destination {
	case "console":
		consoleHandler := slog.NewJSONHandler(os.Stdout, Options)
		L = slog.New(consoleHandler)
	default:
		File, err := os.OpenFile(st.Destination, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return 
		}
		// defer File.Close()
		fileHandler := slog.NewTextHandler(File, Options)
		L = slog.New(fileHandler)
	}
}