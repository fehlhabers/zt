package global

import (
	"sync"

	"github.com/fehlhabers/zt/internal/core/state"
)

var (
	stateKeeper *state.StateKeeper
	once        sync.Once
)

func InitStateKeeper(newStateKeeper *state.StateKeeper) {
	once.Do(func() {
		stateKeeper = newStateKeeper
	})
}

func GetStateKeeper() *state.StateKeeper {
	return stateKeeper
}
