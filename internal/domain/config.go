package domain

import (
	"github.com/fehlhabers/zt/internal/errors"
)

type ZtConfig struct {
	ActiveTeam string                `json:"active_team,omitempty"`
	Teams      map[string]TeamConfig `json:"teams,omitempty"`
}

func NewZtConfig(teamName string, team TeamConfig) *ZtConfig {
	cfg := &ZtConfig{
		ActiveTeam: teamName,
		Teams:      make(map[string]TeamConfig),
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

	return true
}

func (z *ZtConfig) ActiveTeamConfig() (teamName string, teamConfig *TeamConfig, err error) {

	if team, ok := z.Teams[z.ActiveTeam]; ok {
		return z.ActiveTeam, &team, nil
	}

	return "", nil, errors.NoTeamSet
}

type TeamConfig struct {
	SessionDurMins int    `json:"session_dur_mins,omitempty"`
	MainBranch     string `json:"main_branch,omitempty"`
}
