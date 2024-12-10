package config

import (
	"os"
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
	sut := NewConfigRepo(os.TempDir())
	t.Run("add team happy case", func(t *testing.T) {

		err := sut.AddTeam("testTeam", team)
		assert.NoError(t, err)

		actualZtConfig, err := sut.readConfig()
		assert.NoError(t, err)

		assert.Equal(t, actualZtConfig.Teams["testTeam"], team)
	})

	t.Run("switch team not found", func(t *testing.T) {
		err := sut.SwitchTeam("testTeam2")
		assert.Error(t, errors.TeamNotFound, err)
	})

	t.Run("switch team - happy case", func(t *testing.T) {
		err := sut.AddTeam("testTeam2", team)
		assert.NoError(t, err)

		err = sut.SwitchTeam("testTeam2")
		assert.NoError(t, err)

		actualCfg, err := sut.readConfig()
		assert.NoError(t, err)

		assert.Equal(t, "testTeam2", actualCfg.ActiveTeam)
	})
}

func TestReadConfig(t *testing.T) {
	sut := NewConfigRepo(os.TempDir())
	t.Run("error if file not found", func(t *testing.T) {

		_, err := sut.readConfig()

		assert.Error(t, errors.NoZtConfigFound, err)
	})
}
