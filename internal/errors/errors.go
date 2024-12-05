package errors

import "errors"

var (
	TeamNotFound    = errors.New("team does not exist")
	NoTeamSet       = errors.New("no team has been configured")
	NoZtConfigFound = errors.New("no zt configuration found")
)
