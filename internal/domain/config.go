package domain

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

	if _, ok := z.Teams[z.ActiveTeam]; !ok {
		return false
	}

	return true
}

func (z *ZtConfig) ActiveTeamConfig() *ZtTeamConfig {

	if team, ok := z.Teams[z.ActiveTeam]; ok {
		return &ZtTeamConfig{
			Name:           z.ActiveTeam,
			SessionDurMins: team.SessionDurMins,
			MainBranch:     team.MainBranch,
		}
	}

	return nil
}

type TeamConfig struct {
	SessionDurMins int    `json:"session_dur_mins,omitempty"`
	MainBranch     string `json:"main_branch,omitempty"`
}

type ZtTeamConfig struct {
	Name           string
	SessionDurMins int
	MainBranch     string
}
