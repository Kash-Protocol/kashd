package main

import (
	"fmt"
	"os"

	"github.com/Kash-Protocol/kashd/infrastructure/logger"
	"github.com/Kash-Protocol/kashd/stability-tests/common"
	"github.com/Kash-Protocol/kashd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("NTSN")
	spawn      = panics.GoroutineWrapperFunc(log)
)

func initLog(logFile, errLogFile string) {
	level := logger.LevelInfo
	if activeConfig().LogLevel != "" {
		var ok bool
		level, ok = logger.LevelFromString(activeConfig().LogLevel)
		if !ok {
			fmt.Fprintf(os.Stderr, "Log level %s doesn't exists", activeConfig().LogLevel)
			os.Exit(1)
		}
	}
	log.SetLevel(level)
	common.InitBackend(backendLog, logFile, errLogFile)
}
