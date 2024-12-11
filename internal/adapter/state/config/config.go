package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/domain"
	"github.com/fehlhabers/zt/internal/errors"
)

type ZtConfig struct {
	User       string                        `json:"user,omitempty"`
	ActiveTeam string                        `json:"active_team,omitempty"`
	Teams      map[string]*domain.TeamConfig `json:"teams,omitempty"`
}

type ConfigRepo struct {
	statePath string
}

func NewConfigRepo(statePath string) *ConfigRepo {
	return &ConfigRepo{
		statePath: statePath,
	}
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

func (c *ConfigRepo) AddTeam(teamName string, team *domain.TeamConfig) error {
	ztConfig, err := c.readConfig()
	if err != nil || !ztConfig.Valid() {
		log.Info("No configuration found. Creating new config 👌")
		ztConfig = NewZtConfig(teamName, team)
		c.writeConfig(ztConfig)
	} else {
		ztConfig.Teams[teamName] = team
		c.writeConfig(ztConfig)
	}

	return nil
}

func (c *ConfigRepo) ListTeams() map[string]*domain.TeamConfig {
	cfg, err := c.readConfig()
	if err != nil {
		return make(map[string]*domain.TeamConfig)
	}
	return cfg.Teams
}

func (c *ConfigRepo) SwitchTeam(teamName string) error {
	cfg, err := c.readConfig()
	if err != nil {
		return fmt.Errorf("unable to switch team - error: %w", err)
	}

	if _, ok := cfg.Teams[teamName]; !ok {
		return errors.TeamNotFound
	}

	cfg.ActiveTeam = teamName
	if err := c.writeConfig(cfg); err != nil {
		return fmt.Errorf("unable to write config. error=%w", err)
	}
	return nil
}

func (c *ConfigRepo) writeConfig(ztCfg *ZtConfig) error {
	if _, err := os.ReadDir(c.statePath); err != nil {
		os.MkdirAll(c.statePath, os.ModePerm)
	}

	cfg, err := os.OpenFile(c.getCfgFile(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
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

func (c *ConfigRepo) getCfgFile() string {
	return fmt.Sprintf("%s/%s", c.statePath, "config")
}

func (c *ConfigRepo) Reload(state *domain.ZtState) error {
	cfg, err := c.readConfig()
	if err != nil {
		return err
	}

	state.TeamName = cfg.ActiveTeam
	state.AllTeams = cfg.Teams
	state.User = cfg.User

	if team, ok := cfg.Teams[cfg.ActiveTeam]; ok {
		state.TeamConfig = team
	}
	return nil
}

func (c *ConfigRepo) readConfig() (*ZtConfig, error) {
	configReader, err := os.Open(c.getCfgFile())
	if err != nil {
		return nil, errors.NoZtConfigFound
	}

	defer configReader.Close()

	ztConfig := &ZtConfig{}

	if err := json.NewDecoder(configReader).Decode(ztConfig); err != nil {
		return nil, err
	}

	if ztConfig.User == "" {

		var stderr, stdoutBuf bytes.Buffer
		cmd := exec.Command("git", "config", "--get", "user.name")
		cmd.Stderr = &stderr
		cmd.Stdout = &stdoutBuf

		err = cmd.Run()
		if err != nil {
			log.Warn("Tried to resolve user from git, but failed", "error", stderr.String())
		} else {
			ztConfig.User = strings.Trim(stdoutBuf.String(), "\n")
			c.writeConfig(ztConfig)
		}
	}

	return ztConfig, nil
}
