package assert

import (
	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/adapter/state"
)

func Start() error {
	team, err := state.Storer.GetActiveTeam()
	if team == nil || err != nil {
		log.Fatal("no active team configured. Please run 'zt team init'")
	}

	return nil
}
