package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/domain"
	"github.com/fehlhabers/zt/internal/errors"
)

type ZtConfig struct {
	ActiveTeam string                        `json:"active_team,omitempty"`
	Teams      map[string]*domain.TeamConfig `json:"teams,omitempty"`
}

type ConfigRepo struct {
	statePath string
}

func NewZtConfig(teamName string, team *domain.TeamConfig) *ZtConfig {
	cfg := &ZtConfig{
		ActiveTeam: teamName,
		Teams:      make(map[string]*domain.TeamConfig),
	}

	cfg.Teams[teamName] = team

	return cfg
}

func (z *ZtConfig) Valid() bool {
	if z.ActiveTeam == "" {
		return false
	}

	if z.Teams == nil {
		return false
	}

	if _, ok := z.Teams[z.ActiveTeam]; !ok {
		return false
	}

	return true
}

func AddTeam(teamName string, team *domain.TeamConfig) error {
	ztConfig, err := readConfig()
	if err != nil || !ztConfig.Valid() {
		log.Info("No configuration found. Creating new config 👌")
		ztConfig = NewZtConfig(teamName, team)
	} else {
		ztConfig.Teams[teamName] = team
	}

	if err := writeConfig(ztConfig); err != nil {
		return fmt.Errorf("unable to write config. error=%w", err)
	}

	return nil
}

func ListTeams() map[string]*domain.TeamConfig {
	cfg, err := readConfig()
	if err != nil {
		return make(map[string]*domain.TeamConfig)
	}
	return cfg.Teams
}

func SwitchTeam(teamName string) error {
	cfg, err := readConfig()
	if err != nil {
		return fmt.Errorf("unable to switch team - error: %w", err)
	}

	if _, ok := cfg.Teams[teamName]; !ok {
		return errors.TeamNotFound
	}

	cfg.ActiveTeam = teamName
	if err := writeConfig(cfg); err != nil {
		return fmt.Errorf("unable to write config. error=%w", err)
	}
	return nil
}

// Internal writer to be used by exported methods
func writeConfig(ztCfg *ZtConfig) error {
	if _, err := os.ReadDir(state.GetStatePath()); err != nil {
		os.MkdirAll(state.GetStatePath(), os.ModePerm)
	}

	cfg, err := os.OpenFile(getCfgFile(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Error("failed to write to file")
		return err
	}
	defer cfg.Close()

	prettyJson, err := json.MarshalIndent(ztCfg, "", "  ")
	if err != nil {
		return err
	}

	if _, err := cfg.Write(prettyJson); err != nil {
		return err
	}

	return nil
}

func getCfgFile() string {
	return fmt.Sprintf("%s/%s", state.GetStatePath(), "config")
}

func Reload(state *domain.ZtState) error {
	cfg, err := readConfig()
	if err != nil {
		return err
	}

	state.TeamName = cfg.ActiveTeam
	if team, ok := cfg.Teams[cfg.ActiveTeam]; ok {
		state.TeamConfig = team
	}
	return nil
}

func readConfig() (*ZtConfig, error) {
	configReader, err := os.Open(getCfgFile())
	if err != nil {
		return nil, errors.NoZtConfigFound
	}

	defer configReader.Close()

	ztConfig := &ZtConfig{}

	if err := json.NewDecoder(configReader).Decode(ztConfig); err != nil {
		return nil, err
	}

	return ztConfig, nil
}
