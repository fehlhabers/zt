package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/cmd"
	"github.com/fehlhabers/zt/internal/core/state"
	"github.com/fehlhabers/zt/internal/global"
)

const (
	stateDir = ".local/state/zt"
)

var (
	Version string = "0.1.0"
)

func main() {
	log.SetReportTimestamp(false)

	stateKeeper := state.NewStateKeeper(getDefaultStateDir())
	global.InitStateKeeper(stateKeeper)

	err := cmd.NewRoot().Execute()
	if err != nil {
		os.Exit(1)
	}

}

func getDefaultStateDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Warn("Unable to get home directory!")
	}
	return fmt.Sprintf("%s/%s", home, stateDir)
}
