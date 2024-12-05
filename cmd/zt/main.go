package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/adapter/state"
	"github.com/fehlhabers/zt/internal/cmd"
	"github.com/fehlhabers/zt/internal/core/config"
)

var Version string = "0.1.0"

func main() {
	log.SetReportTimestamp(false)

	state.Storer = state.NewZtreamStorer(config.GetCfgFileDir())
	config.ReloadConfig()

	err := cmd.NewRoot().Execute()
	if err != nil {
		os.Exit(1)
	}

}
