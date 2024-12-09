package domain

import "time"

type ZtState struct {
	User          string
	TeamName      string
	TeamConfig    *TeamConfig
	CurrentZtream *Ztream
}

type Ztream struct {
	Name    string
	Started time.Time
	Ends    time.Time
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
	z.Ends = time.Now().Add(time.Duration(sessionMins))
}
