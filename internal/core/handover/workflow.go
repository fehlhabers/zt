package handover

import (
	"fmt"
	"math"
	"time"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/adapter/git"
	"github.com/fehlhabers/zt/internal/adapter/state"
	"github.com/fehlhabers/zt/internal/core/config"
	"github.com/fehlhabers/zt/internal/domain"
	"github.com/fehlhabers/zt/internal/model"
)

func mustGetConfig() *domain.ZtTeamConfig {
	cfg, err := config.ZtConfig()
	if err != nil {
		log.Fatal("Failed to create ztream!", "error", err)
	}

	return cfg.ActiveTeamConfig()
}

func CreateZtream(ztreamName string) {

	teamCfg := mustGetConfig()

	branch, err := git.CurrentBranch()
	if err != nil {
		log.Error("Failed to create ztream!", "error", err)
		return
	}

	if teamCfg.MainBranch != branch {
		log.Warn("It is recommended to start a ztream from main/master")
	}

	now := time.Now().Unix()

	zt := model.Ztream{
		Name:    ztreamName,
		Started: now,
		Ends:    now + int64(teamCfg.SessionDurMins*60),
		Team:    teamCfg.Name,
	}

	if err := state.Storer.StoreZtream(zt); err != nil {
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

	curZt, err := state.Storer.GetActiveZtream()
	if err != nil {
		log.Warn("Unable to fetch current local ztream")
		return
	}

	untilEnd := time.Unix(curZt.Ends, 0).Sub(time.Now())
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

	teamCfg := mustGetConfig()

	now := time.Now().Unix()

	z := model.Ztream{
		Name:    ztreamName,
		Started: now,
		Ends:    now,
		Team:    teamCfg.Name,
	}
	state.Storer.StoreZtream(z)

	log.Info("Joined the ztream! 🙌")
}

func Next() {
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
	if !isActiveZtream() {
		return
	}

	log.Info("Starting ztream...")

	if _, err := git.Pull(); err != nil {
		log.Error("Failed to start handover!", "error", err)
		return
	}

	z, err := state.Storer.GetActiveZtream()
	if err != nil {
		return
	}

	cfg, err := config.ZtConfig()
	if err != nil {
		log.Fatal("Invalid configuration", "error", err)
	}

	teamCfg := cfg.ActiveTeamConfig()
	z.StartSession(teamCfg.SessionDurMins)
	state.Storer.StoreZtream(z)
}

func isActiveZtream() bool {
	branch, err := git.CurrentBranch()
	if err != nil {
		log.Error("Unable to get current branch", "error", err)
		return false
	}

	if ztream, err := state.Storer.GetActiveZtream(); err != nil || ztream.Name != branch {
		log.Error("Join the ztream before starting.", "branch", branch, "active_ztream", ztream.Name)
		return false
	}

	if branch == "master" || branch == "main" {
		log.Error("Cannot start a ztream on main/master!")
		return false
	}

	return true
}
