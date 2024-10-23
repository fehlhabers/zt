package state

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/model"
	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
)

var (
	// global storer since it's a CLI
	Storer *ZtreamStorer

	errMultipleDefaultTeams = errors.New("multiple default teams found")
	errNoDefaultTeam        = errors.New("no default team found")
)

type ZtreamStorer struct {
	db *sqlx.DB
}

func NewZtreamStorer(dbPath string) *ZtreamStorer {

	mustCreateDirectories(dbPath)

	dbFile := fmt.Sprintf("%s/ztream.db", dbPath)
	db, err := sqlx.Open("sqlite", dbFile)
	if err != nil {
		log.Fatal("Not able to create db instance", "path", dbPath, "error", err)
	}

	if _, err := db.Exec(model.CreateTeamsTable); err != nil {
		log.Fatal("Not able to prepare team", "error", err)
	}

	if _, err := db.Exec(model.CreateZtreamsTable); err != nil {
		log.Fatal("Not able to prepare ztream", "error", err)
	}

	if _, err := db.Exec(model.CreateActiveZtreamTable); err != nil {
		log.Fatal("Not able to prepare active ztream", "error", err)
	}

	if _, err := db.Exec(model.CreateActiveTeamTable); err != nil {
		log.Fatal("Not able to prepare active team", "error", err)
	}

	if _, err := db.Exec(model.CreateMembersTable); err != nil {
		log.Fatal("Not able to prepare members", "error", err)
	}

	return &ZtreamStorer{db: db}
}

func (z *ZtreamStorer) GetActiveTeam() (*model.Team, error) {
	var (
		selectActiveTeam = `
		SELECT z.* FROM teams z JOIN active_team USING (name)
		`
	)
	teams := []model.Team{}
	if err := z.db.Select(&teams, selectActiveTeam); err != nil {
		return nil, err
	}

	return &teams[0], nil
}

func (z *ZtreamStorer) GetTeam(name string) (*model.Team, error) {
	team := &model.Team{}
	if err := z.db.Get(team, "SELECT * FROM teams WHERE name = $1", name); err != nil {
		return nil, err
	}

	return team, nil
}

func (z *ZtreamStorer) GetTeams() ([]*model.Team, error) {
	teams := []*model.Team{}
	if err := z.db.Select(teams, "SELECT * FROM teams"); err != nil {
		return nil, err
	}

	return teams, nil
}

func (z *ZtreamStorer) CreateTeam(team model.Team) error {
	var (
		insertTeam = `
		INSERT INTO teams (
			name
		)
		VALUES (
			:name
		)
		`
		setActiveTeam = `
		INSERT INTO active_team (
			only,
			name
			) VALUES (
			    1,
				:name
			)
		ON CONFLICT (only) DO UPDATE SET
			name = :name
		`
	)

	if _, err := z.db.NamedExec(insertTeam, &team); err != nil {
		return err
	}

	if _, err := z.db.NamedExec(setActiveTeam, &team); err != nil {
		return err
	}

	return nil
}

func (z *ZtreamStorer) GetActiveZtream() (model.Ztream, error) {
	var zt model.Ztream
	var (
		selectActiveZtream = `
		SELECT z.* FROM ztreams z JOIN active_ztream USING (name)
		`
	)

	if err := z.db.Get(&zt, selectActiveZtream); err != nil {
		return zt, err
	}
	return zt, nil
}

// StoreZtream implements model.Storer.
func (z *ZtreamStorer) StoreZtream(zt model.Ztream) error {
	var (
		setActiveZstream = `
		INSERT INTO active_ztream (
			only,
			name
			) VALUES (
			    1,
				:name
			)
		ON CONFLICT (only) DO UPDATE SET
			name = :name
		`
	)
	if _, err := z.db.NamedExec(model.InsertZtream, &zt); err != nil {
		return err
	}

	if _, err := z.db.NamedExec(setActiveZstream, &zt); err != nil {
		return err
	}

	return nil
}

func mustCreateDirectories(path string) {
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		log.Fatal("Not able to create directories to store ztream data", "path", path, "error", err)
	}
}
