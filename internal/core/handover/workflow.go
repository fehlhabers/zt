package handover

import (
	"math"
	"slices"
	"time"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/adapter/git"
	"github.com/fehlhabers/zt/internal/adapter/state"
	"github.com/fehlhabers/zt/internal/errors"
	"github.com/fehlhabers/zt/internal/model"
)

var (
	mainBranches = []string{
		"main",
		"master",
	}
)

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

func CreateZtream(ztreamName string) {
	team, err := state.Storer.GetActiveTeam()
	if err != nil {
		log.Fatal(errors.GuiTeamNotSet, "error", err)
	}

	branch, err := git.CurrentBranch()
	if err != nil {
		log.Error("Failed to start handover!", "error", err)
		return
	}

	if slices.Contains(mainBranches, branch) {
		log.Warn("It is recommended to start a ztream from main/master")
	}

	now := time.Now().Unix()

	zt := model.Ztream{
		Name:    ztreamName,
		Started: now,
		Ends:    now + 60*15,
		Team:    team.Name,
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

	log.Info("Started ztream!", "ztream", curZt.Name, "mins left", minsUntil)
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

	if branch == "master" || branch == "main" {
		return false
	}

	return true
}
