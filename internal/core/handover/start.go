package handover

import (
	"fmt"
	"strings"

	"github.com/fehlhabers/st/internal/core/git"
)

const (
	activeSesssionPrefix = "st"
)

func CreateSession(sessionName string) {
	branch, err := git.CurrentBranch()
	if err != nil {
		fmt.Printf("Failed to start handover!\n%s\n", err)
		return
	}

	if branch != "master" && branch != "main" {
		fmt.Println("Warning! It is recommended to use Stream Team from the trunk")
	}
	createSession(sessionName)
}

func createSession(sessionName string) {
	session := fmt.Sprintf("%s/%s", activeSesssionPrefix, sessionName)
	if _, err := git.CreateBranch(session); err != nil {
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

	branchParts := strings.Split(branch, "/")
	if len(branchParts) < 2 {
		return false
	}

	if branchParts[0] == activeSesssionPrefix {
		return true
	}
	return false
}
