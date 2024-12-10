package domain

type TeamConfig struct {
	SessionDurMins int    `json:"session_dur_mins,omitempty"`
	MainBranch     string `json:"main_branch,omitempty"`
	MergeStrategy  string `json:"merge_strategy,omitempty"`
}

var (
	MergeStrategyGithubPr = "github-pr"
)
