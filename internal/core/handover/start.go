package handover

import (
	"fmt"
	"slices"

	"github.com/fehlhabers/st/internal/core/git"
)

var (
	mainBranches = []string{
		"main",
		"master",
	}
)

func JoinSession(sessionName string) {
	_, err := git.Pull()
	if err != nil {
		fmt.Printf("Failed to join stream!\n%s\n", err)
		return
	}
}

func CreateSession(sessionName string) {
	branch, err := git.CurrentBranch()
	if err != nil {
		fmt.Printf("Failed to start handover!\n%s\n", err)
		return
	}

	if slices.Contains(mainBranches, branch) {
		fmt.Println("Warning! It is recommended to use Stream Team from the trunk")
	}
	createSession(sessionName)
}

func createSession(sessionName string) {
	if _, err := git.CreateBranch(sessionName); err != nil {
		fmt.Printf("Failed to start handover!\n%s\n", err)
		return
	}
}

func Next() {
	if isActiveSession() {
		fmt.Println("Handing over...")
		git.AddAll()
		git.Commit("Stream Team - Next")
		git.Push()
		fmt.Println("Handover done!")
	} else {
		fmt.Println("No active session found! Make sure you are in the right branch")
	}
}

func Start() {
	if isActiveSession() {
		fmt.Println("Starting session...")
		if _, err := git.Pull(); err != nil {
			fmt.Printf("Failed to start handover!\n%s\n", err)
			return
		}
	} else {
		fmt.Println("No active session found! Create a session before starting")
	}
}

func isActiveSession() bool {
	branch, err := git.CurrentBranch()
	if err != nil {
		return false
	}

	if branch == "master" || branch == "main" {
		return false
	}

	return true
}
