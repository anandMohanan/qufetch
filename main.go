package main

import (
	"github.com/charmbracelet/log"
	"github.com/anandMohanan/qufetch/cmd"
	"github.com/anandMohanan/qufetch/config"
	"github.com/anandMohanan/qufetch/logger"
	"github.com/samber/lo"
	"os"
)

func handlePanic() {
	if err := recover(); err != nil {
		log.Error("crashed", "err", err)
		os.Exit(1)
	}
}

func main() {
	defer handlePanic()

	// prepare config and logs
	lo.Must0(config.Init())
	lo.Must0(logger.Init())

	// run the app
	cmd.Execute()
}
