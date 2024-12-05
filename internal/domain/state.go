package domain

import "time"

type ZtState struct {
	User          string
	TeamName      string
	TeamConfig    TeamConfig
	CurrentZtream Ztream
}

type Ztream struct {
	Name    string `db:"name"`
	Started int64  `db:"started"`
	Ends    int64  `db:"ends"`
}

func (z *Ztream) StartSession(sessionMins int) {
	startTime := time.Now()
	endTime := startTime.Add(time.Duration(sessionMins))

	_ = endTime
	z.Started = startTime.Unix()
	z.Ends = endTime.Unix()
}
