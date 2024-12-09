package state

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

const (
	stateDir = ".local/state/zt"
)

func GetDefaultStateDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Warn("Unable to get home directory!")
	}
	return fmt.Sprintf("%s/%s", home, stateDir)
}
