package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/domain"
	"github.com/fehlhabers/zt/internal/errors"
)

const (
	CfgHomeFileDir = ".local/state/zt"
)

var (
	zt *domain.ZtConfig
)

func ZtConfig() (domain.ZtConfig, error) {
	if !zt.Valid() {
		return domain.ZtConfig{}, errors.NoZtConfigFound
	}
	return *zt, nil
}

func AddTeam(teamName string, team domain.TeamConfig) error {
	ztConfig, err := readConfig()
	if err != nil || !ztConfig.Valid() {
		log.Info("No configuration found. Creating new config 👌")
		ztConfig = domain.NewZtConfig(teamName, team)
	} else {
		ztConfig.Teams[teamName] = team
	}

	if err := writeConfig(ztConfig); err != nil {
		return fmt.Errorf("unable to write config. error=%w", err)
	}

	ReloadConfig()
	return nil
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
	ReloadConfig()
	return nil
}

// Internal writer to be used by exported methods
func writeConfig(ztCfg *domain.ZtConfig) error {
	if _, err := os.ReadDir(GetCfgFileDir()); err != nil {
		os.MkdirAll(GetCfgFileDir(), os.ModePerm)
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

func GetCfgFileDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Warn("Unable to get home directory!")
	}
	return fmt.Sprintf("%s/%s", home, CfgHomeFileDir)
}

func getCfgFile() string {
	return fmt.Sprintf("%s/%s", GetCfgFileDir(), "config")
}

func ReloadConfig() error {
	cfg, err := readConfig()
	if err != nil {
		return err
	}
	zt = cfg

	return nil
}

func readConfig() (*domain.ZtConfig, error) {
	configReader, err := os.Open(getCfgFile())
	if err != nil {
		return nil, errors.NoZtConfigFound
	}

	defer configReader.Close()

	ztConfig := &domain.ZtConfig{}

	if err := json.NewDecoder(configReader).Decode(ztConfig); err != nil {
		return nil, err
	}

	return ztConfig, nil
}
