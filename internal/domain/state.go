package domain

import (
	"errors"
	"time"
)

type ZtState struct {
	User       string
	TeamName   string
	TeamConfig *TeamConfig
	AllTeams   map[string]*TeamConfig
	CurZtream  *Ztream
	AllZtreams []*Ztream
}

type Ztream struct {
	Name     string
	Metadata string
	Started  time.Time
	Ends     time.Time
}

func NewZtream(name string, cfg *TeamConfig) *Ztream {
	z := &Ztream{
		Name: name,
	}
	z.StartSession(cfg.SessionDurMins)
	return z
}

func (z *Ztream) StartSession(sessionMins int) {
	z.Started = time.Now()
	z.Ends = time.Now().Add(time.Minute * time.Duration(sessionMins))
}

func (z *ZtState) HasActiveZtream() bool {
	return z.CurZtream != nil
}

func (z *ZtState) Validate() error {
	if z.TeamConfig == nil {
		return errors.New("no team configuration found. Run 'zt team configure'.")
	}

	if z.TeamName == "" {
		return errors.New("no team configuration found. Run 'zt team configure'.")
	}

	if z.User == "" {
		return errors.New("no user configured. Run zt in your git repo to use git user")
	}

	return nil
}
