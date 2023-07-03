package main

import (
	"fmt"
	"os"

	"github.com/open-binance/logger"
)

func main() {
	cfgLog := logger.NewDefaultCfg()
	if err := logger.Init(cfgLog); err != nil {
		fmt.Printf("msg=failed to init logger||err=%s\n", err.Error())
		os.Exit(1)
	}
	defer logger.Sync()

	priceCfg := logger.FileConfig{
		Enable:     true,
		LocalTime:  true,
		Compress:   false,
		MaxSize:    256,
		MaxAge:     7,
		MaxBackups: 32,
		Level:      "info",
		Filename:   "logs/price.txt", // the file to write logs to
		TimeKey:    "",
		CallerKey:  "",
		LevelKey:   "",
	}
	priceLogger, err := logger.NewFileLogger(priceCfg)
	if err != nil {
		logger.Errorf("msg=failed to new file logger")
		return
	}
	priceLogger.Info("hello world")

	logger.Infof("msg=succeed to init logger")
}
