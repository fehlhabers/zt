package handover

import (
	"slices"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/core/git"
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
	branch, err := git.CurrentBranch()
	if err != nil {
		log.Error("Failed to start handover!", "error", err)
		return
	}

	if slices.Contains(mainBranches, branch) {
		log.Warn("It is recommended to start a ztream from main/master")
	}
	if _, err := git.CreateBranch(ztreamName); err != nil {
		log.Error("Failed to start handover!", "error", err)
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

	if branch == "master" || branch == "main" {
		return false
	}

	return true
}
