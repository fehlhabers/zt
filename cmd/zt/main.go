package main

import (
	"os"

	"github.com/charmbracelet/log"
	zt "github.com/fehlhabers/zt/internal/adapter/state"
	"github.com/fehlhabers/zt/internal/cmd"
	"github.com/fehlhabers/zt/internal/core/state"
)

var Version string = "0.1.0"

func main() {
	log.SetReportTimestamp(false)

	statePath := zt.GetDefaultStateDir()
	state.Init(statePath)
	err := cmd.NewRoot().Execute()
	if err != nil {
		os.Exit(1)
	}

}
