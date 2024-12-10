package admin

import (
	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/global"
)

func ListZtreams() {

	for _, z := range global.GetStateKeeper().GetState().AllZtreams {
		log.Info("Ztream", "name", z.Name)
	}
}
