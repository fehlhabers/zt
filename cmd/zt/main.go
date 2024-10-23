package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/adapter/state"
	"github.com/fehlhabers/zt/internal/cmd"
)

var Version string = "0.1.0"

func main() {
	log.SetReportTimestamp(false)

	homeDir, _ := os.UserHomeDir()
	dbPath := fmt.Sprintf("%s/.local/state", homeDir)
	storer := state.NewZtreamStorer(dbPath)

	err := cmd.NewRoot(storer).Execute()
	if err != nil {
		os.Exit(1)
	}

}
