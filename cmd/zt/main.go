package main

import (
	"os"
	"time"

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

	if ztream, err := state.Storer.GetActiveZtream(); err != nil {
		endTime := time.Unix(ztream.Ends, 0)
		endsIn := endTime.Sub(time.Now())

		log.Info("Current ztream: ", "name", ztream.Name, "ends at", endsIn.String()+" mins")
	}

	err := cmd.NewRoot().Execute()
	if err != nil {
		os.Exit(1)
	}

}
