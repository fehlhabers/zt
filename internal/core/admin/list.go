package admin

import "github.com/fehlhabers/zt/internal/adapter/state"

func ListZtreams() {
	state.Storer.GetActiveZtream()
}
