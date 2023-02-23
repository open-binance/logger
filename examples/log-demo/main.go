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

	logger.Infof("msg=succeed to init logger")
}
