package handover

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"os/exec"
	"time"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/adapter/git"
	"github.com/fehlhabers/zt/internal/core/timer"
	"github.com/fehlhabers/zt/internal/domain"
	"github.com/fehlhabers/zt/internal/global"
)

func CreateZtream(ztreamName string, metadata string) {

	ztState := global.GetStateKeeper().GetState()

	if err := ztState.Validate(); err != nil {
		log.Error("Missing configuration", "error", err)
		return
	}

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

	if !isOnZtreamBranch() {
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
	var (
		zt         = global.GetStateKeeper().GetState()
		ztreamName = zt.CurZtream.Name
	)

	_, err := git.Fetch()
	if err != nil {
		log.Fatal("Failed to fetch from remote")
	}

	branch, err := git.CurrentBranch()
	if err != nil {
		log.Fatal("Failed to get current branch")
	}

	if !zt.HasActiveZtream() {
		log.Fatal("No active ztream found")
	}

	if zt.CurZtream.Name != branch {
		log.Info("Not on ztream. Trying to switch/create", "ztream", zt.CurZtream.Name)
		_, err := git.SwitchBranch(zt.CurZtream.Name)
		if err != nil {
			fmt.Printf("git checkout -b %s\n", ztreamName)
			if _, err := git.CreateBranch(ztreamName); err != nil {
				log.Fatal("Failed to start handover!", "error", err)
				return
			}

			fmt.Printf("git push --set-upstream orgin %s\n", ztreamName)
			if _, err := git.PushSetOrigin(ztreamName); err != nil {
				log.Fatal("Failed to start handover!", "error", err)
				return
			}
			log.Info("Created branch!", "ztream", zt.CurZtream.Name)
		}
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

func Merge() {
	ztState := global.GetStateKeeper().GetState()
	if !ztState.HasActiveZtream() {
		log.Fatal("No active ztream found")
	}

	if !isOnZtreamBranch(ztState) {
		log.Fatal("No active ztream found")
	}

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("gh", "pr", "create", "--title", ztState.CurZtream.Name, "--body", ztState.CurZtream.Metadata)

	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err := cmd.Run()

	log.Info(stdout.String())

	if err != nil {
		log.Error("Unable to create pull request", "error", stderr.String())
	}
}

func isOnZtreamBranch(zt domain.ZtState) error {
	branch, err := git.CurrentBranch()
	if err != nil {
		return fmt.Errorf("failed to get branch - %w", err)
	}

	if !zt.HasActiveZtream() {
		return errors.New("no active ztream found")
	} else if zt.CurZtream.Name != branch {
		return errors.New("not on ztream branch")
	}

	if branch == "master" || branch == "main" {
		return errors.New("on main/master branch")
	}

	return nil
}

func PrintCurrentZtream() {
	if ztream := global.GetStateKeeper().GetState().CurZtream; ztream != nil {
		endsIn := ztream.Ends.Sub(time.Now()).Round(time.Minute)

		log.Info("Current ztream -", "name", ztream.Name, "ends in", endsIn.String())
	}
}
