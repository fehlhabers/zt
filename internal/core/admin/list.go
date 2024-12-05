package admin

import (
	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/adapter/state"
)

func ListZtreams() {
	for _, z := range state.Storer.GetAllZtreams() {
		log.Info("Ztream", "name", z.Name)
	}
}
