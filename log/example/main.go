package main

import (
	"github.com/feixiao/infra/log/example/a"

	log "github.com/feixiao/infra/log"
)

func main() {
	// defer log4go.Close()
	log.Info("测试")

	log.SetLoggerLevel(log.WARNING)

	log.Info("测试Info")
	log.Warn("测试Warn")

	a.PrintLog()
}
