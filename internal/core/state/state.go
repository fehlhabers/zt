package state

import (
	"github.com/fehlhabers/zt/internal/adapter/state/config"
	"github.com/fehlhabers/zt/internal/adapter/state/ztreams"
	"github.com/fehlhabers/zt/internal/domain"
)

var (
	zt        *domain.ZtState
	Ztreams   *ztreams.ZtreamStorer
	statePath string
)

func Init(statePath string) {
	statePath = statePath
	zt = &domain.ZtState{}
	Ztreams = ztreams.NewZtreamStorer(statePath)
	reloadState()
}

func GetStatePath() string {
	return statePath
}

func GetState() domain.ZtState {
	reloadState()
	return *zt
}

func reloadState() {
	config.Reload(zt)
	Ztreams.Reload(zt)
}
