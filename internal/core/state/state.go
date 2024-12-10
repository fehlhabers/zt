package state

import (
	"github.com/fehlhabers/zt/internal/adapter/state/config"
	"github.com/fehlhabers/zt/internal/adapter/state/ztreams"
	"github.com/fehlhabers/zt/internal/domain"
)

type StateKeeper struct {
	state      *domain.ZtState
	ztreamRepo *ztreams.ZtreamRepo
	configRepo *config.ConfigRepo
}

func NewStateKeeper(statePath string) *StateKeeper {
	return &StateKeeper{
		state:      &domain.ZtState{},
		ztreamRepo: ztreams.NewZtreamStorer(statePath),
		configRepo: config.NewConfigRepo(statePath),
	}
}

func (s *StateKeeper) GetConfigRepo() *config.ConfigRepo {
	return s.configRepo
}

func (s *StateKeeper) GetZtreamRepo() *ztreams.ZtreamRepo {
	return s.ztreamRepo
}

func (s *StateKeeper) GetState() domain.ZtState {
	s.reloadState()
	return *s.state
}

func (s *StateKeeper) reloadState() {
	s.configRepo.Reload(s.state)
	s.ztreamRepo.Reload(s.state)
}
