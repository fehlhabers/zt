package handover

import (
	"math"
	"time"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/adapter/git"
	"github.com/fehlhabers/zt/internal/adapter/state"
	"github.com/fehlhabers/zt/internal/core/config"
	"github.com/fehlhabers/zt/internal/model"
)

var (
	mainBranches = []string{
		"main",
		"master",
	}
)

func CreateZtream(ztreamName string) {
	cfg, err := config.ZtConfig()
	if err != nil {
		log.Error("Failed to create ztream!", "error", err)
	}

	teamName, teamCfg, err := cfg.ActiveTeamConfig()
	if err != nil {
		log.Error("Failed to create ztream!", "error", err)
		return
	}

	branch, err := git.CurrentBranch()
	if err != nil {
		log.Error("Failed to create ztream!", "error", err)
		return
	}

	if teamCfg.MainBranch == branch {
		log.Warn("It is recommended to start a ztream from main/master")
	}

	now := time.Now().Unix()

	zt := model.Ztream{
		Name:    ztreamName,
		Started: now,
		Ends:    now + int64(teamCfg.SessionDurMins*60),
		Team:    teamName,
	}

	if err := state.Storer.StoreZtream(zt); err != nil {
		log.Fatal("Unable to save new ztream", "error", err)
	}

	if _, err := git.CreateBranch(ztreamName); err != nil {
		log.Error("Failed to start handover!", "error", err)
		return
	}

	curZt, err := state.Storer.GetActiveZtream()
	if err != nil {
		log.Warn("Unable to fetch current local ztream")
		return
	}

	untilEnd := time.Unix(curZt.Ends, 0).Sub(time.Now())
	minsUntil := math.Trunc(untilEnd.Minutes())
	log.Info("Started ztream", "name", curZt.Name, "mins left", minsUntil)
}

func JoinZtream(ztreamName string) {
	_, err := git.Fetch()
	if err != nil {
		log.Warn("Failed updating from remote", "error", err)
		return
	}

	_, err = git.SwitchBranch(ztreamName)
	if err != nil {
		log.Warn("Could not join Ztream. Does it exist?", "error", err)
		return
	}

	_, err = git.Pull()
	if err != nil {
		log.Warn("Unable to pull latest changes from remote", "error", err)
		return
	}
}

func Next() {
	if isActiveZtream() {
		log.Info("Handing over...")
		git.AddAll()
		git.Commit("zt handover")
		git.Push()
		log.Info("Handover done!")
	} else {
		log.Error("No active ztream found! Make sure you are in the right branch")
	}
}

func Start() {
	if isActiveZtream() {
		log.Info("Starting ztream...")
		if _, err := git.Pull(); err != nil {
			log.Error("Failed to start handover!", "error", err)
			return
		}
	} else {
		log.Error("No active ztream found! Create a ztream before starting")
	}
}

func isActiveZtream() bool {
	branch, err := git.CurrentBranch()
	if err != nil {
		return false
	}

	if ztream, err := state.Storer.GetActiveZtream(); err != nil || ztream.Name != branch {
		log.Error("Join the ztream before starting.", "branch", branch, "active ztream", ztream.Name)
		return false
	}

	if branch == "master" || branch == "main" {
		return false
	}

	return true
}
