package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/cmd"
)

func main() {
	log.SetReportTimestamp(false)
	err := cmd.NewRoot().Execute()
	if err != nil {
		os.Exit(1)
	}

}
