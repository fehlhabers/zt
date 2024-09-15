package handover

import (
	"fmt"
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
	_, err := git.Pull()
	if err != nil {
		log.Warn("Failed to join stream!\n%s\n", err)
		return
	}
}

func CreateZtream(ztreamName string) {
	branch, err := git.CurrentBranch()
	if err != nil {
		fmt.Printf("Failed to start handover!\n%s\n", err)
		return
	}

	if slices.Contains(mainBranches, branch) {
		log.Warn("It is recommended to use Stream Team from the trunk")
	}
	if _, err := git.CreateBranch(ztreamName); err != nil {
		fmt.Printf("Failed to start handover!\n%s\n", err)
		return
	}
}

func Next() {
	if isActiveZtream() {
		log.Info("Handing over...")
		git.AddAll()
		git.Commit("Stream Team - Next")
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
			fmt.Printf("Failed to start handover!\n%s\n", err)
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
