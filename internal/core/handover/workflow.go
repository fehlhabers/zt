package handover

import (
	"fmt"
	"math"
	"time"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/adapter/git"
	"github.com/fehlhabers/zt/internal/core/timer"
	"github.com/fehlhabers/zt/internal/domain"
	"github.com/fehlhabers/zt/internal/global"
)

func CreateZtream(ztreamName string) {

	ztState := global.GetStateKeeper().GetState()

	branch, err := git.CurrentBranch()
	if err != nil {
		log.Error("Failed to create ztream!", "error", err)
		return
	}

	if ztState.TeamConfig.MainBranch != branch {
		log.Warn("It is recommended to start a ztream from main/master")
	}

	zt := domain.NewZtream(ztreamName, ztState.TeamConfig)

	if err := global.GetStateKeeper().GetZtreamRepo().StoreZtream(zt); err != nil {
		log.Fatal("Unable to save new ztream", "error", err)
	}

	fmt.Printf("git checkout -b %s\n", ztreamName)
	if _, err := git.CreateBranch(ztreamName); err != nil {
		log.Error("Failed to start handover!", "error", err)
		return
	}

	fmt.Printf("git push --set-upstream orgin %s\n", ztreamName)
	if _, err := git.PushSetOrigin(ztreamName); err != nil {
		log.Error("Failed to start handover!", "error", err)
		return
	}

	curZt := global.GetStateKeeper().GetState().CurZtream

	untilEnd := curZt.Ends.Sub(time.Now())
	minsUntil := math.Round(untilEnd.Minutes())
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

	now := time.Now()

	z := &domain.Ztream{
		Name:    ztreamName,
		Started: now,
		Ends:    now,
	}
	global.GetStateKeeper().GetZtreamRepo().StoreZtream(z)

	log.Info("Joined the ztream! 🙌")
}

func Next() {
	PrintCurrentZtream()

	if !isActiveZtream() {
		log.Error("No active ztream found! Make sure you are in the right branch")
		return
	}

	log.Info("Handing over...")
	git.AddAll()
	git.Commit("zt handover")
	git.Push()
	log.Info("Handover done! ✅")
}

func Start() {
	ztState := global.GetStateKeeper().GetState()

	if !isActiveZtream() {
		return
	}

	log.Info("Starting ztream...")

	if _, err := git.Pull(); err != nil {
		log.Error("Failed to start handover!", "error", err)
		return
	}

	z := ztState.CurZtream
	teamCfg := ztState.TeamConfig
	z.StartSession(teamCfg.SessionDurMins)
	global.GetStateKeeper().GetZtreamRepo().StoreZtream(z)
	timer.Start(&ztState)
	PrintCurrentZtream()
}

func isActiveZtream() bool {
	branch, err := git.CurrentBranch()
	if err != nil {
		log.Error("Unable to get current branch", "error", err)
		return false
	}
	ztream := global.GetStateKeeper().GetState().CurZtream
	if ztream == nil {
		log.Error("No active ztream. Join or create a ztream before starting")
		return false
	} else if ztream.Name != branch {
		log.Error("Join the ztream before starting.", "branch", branch, "active_ztream", ztream.Name)
		return false
	}

	if branch == "master" || branch == "main" {
		log.Error("Cannot start a ztream on main/master!")
		return false
	}

	return true
}

func PrintCurrentZtream() {
	if ztream := global.GetStateKeeper().GetState().CurZtream; ztream != nil {
		endsIn := ztream.Ends.Sub(time.Now()).Round(time.Minute)

		log.Info("Current ztream -", "name", ztream.Name, "ends in", endsIn.String())
	}
}
