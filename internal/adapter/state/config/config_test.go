package config

import (
	"testing"

	"github.com/fehlhabers/zt/internal/domain"
	"github.com/fehlhabers/zt/internal/errors"
	"github.com/stretchr/testify/assert"
)

func TestAlterTeam(t *testing.T) {
	team := &domain.TeamConfig{
		SessionDurMins: 10,
		MainBranch:     "main",
	}
	t.Run("add team happy case", func(t *testing.T) {

		err := AddTeam("testTeam", team)
		assert.NoError(t, err)

		actualZtConfig, err := readConfig()
		assert.NoError(t, err)

		assert.Equal(t, actualZtConfig.Teams["testTeam"], team)
	})

	t.Run("switch team not found", func(t *testing.T) {
		err := SwitchTeam("testTeam2")
		assert.Error(t, errors.TeamNotFound, err)
	})

	t.Run("switch team - happy case", func(t *testing.T) {
		err := AddTeam("testTeam2", team)
		assert.NoError(t, err)

		err = SwitchTeam("testTeam2")
		assert.NoError(t, err)

		actualCfg, err := readConfig()
		assert.NoError(t, err)

		assert.Equal(t, "testTeam2", actualCfg.ActiveTeam)
	})
}

func TestReadConfig(t *testing.T) {
	t.Run("error if file not found", func(t *testing.T) {

		_, err := readConfig()

		assert.Error(t, errors.NoZtConfigFound, err)
	})
}
