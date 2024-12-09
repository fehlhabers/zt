package admin

import (
	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/core/state"
)

func ListZtreams() {

	for _, z := range state.Ztreams.GetAllZtreams() {
		log.Info("Ztream", "name", z.Name)
	}
}
