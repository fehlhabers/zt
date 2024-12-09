package ztreams

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/adapter/state"
	"github.com/fehlhabers/zt/internal/domain"
	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
)

var (
	errMultipleDefaultTeams = errors.New("multiple default teams found")
	errNoDefaultTeam        = errors.New("no default team found")
)

type ZtreamStorer struct {
	db *sqlx.DB
}

func NewZtreamStorer(dbPath string) *ZtreamStorer {

	mustCreateDirectories(dbPath)

	dbFile := fmt.Sprintf("%s/state.db", dbPath)
	db, err := sqlx.Open("sqlite", dbFile)
	if err != nil {
		log.Fatal("Not able to create db instance", "path", dbPath, "error", err)
	}

	if _, err := db.Exec(state.CreateZtreamsTable); err != nil {
		log.Fatal("Not able to prepare ztream", "error", err)
	}

	if _, err := db.Exec(state.CreateActivesTable); err != nil {
		log.Fatal("Not able to prepare actives", "error", err)
	}

	return &ZtreamStorer{db: db}
}

func (z *ZtreamStorer) Reload(zt *domain.ZtState) {

	ztream, err := z.GetActiveZtream()
	if err != nil {
		return
	}

	zt.CurrentZtream = ztream
}

func (z *ZtreamStorer) GetAllZtreams() []domain.Ztream {
	var (
		selectAllZtreams = `
		SELECT z.* FROM ztreams z
		`
	)
	rows, err := z.db.Query(selectAllZtreams)
	if err != nil {
		return nil
	}

	var ztreams []domain.Ztream
	for rows.Next() {
		var z domain.Ztream

		err := rows.Scan(&z.Name, &z.Started, &z.Ends)
		if err != nil {
			log.Fatal(err)
		}

		ztreams = append(ztreams, z)
	}

	return ztreams
}

func (z *ZtreamStorer) GetActiveZtream() (*domain.Ztream, error) {
	var zt state.ZtreamDb
	var (
		selectActiveZtream = `
		SELECT z.* FROM ztreams z INNER JOIN actives a ON (a.reference = z.name) WHERE a.type = 'ztream'
		`
	)

	if err := z.db.Get(&zt, selectActiveZtream); err != nil {
		return nil, err
	}
	return zt.ToZtream(), nil
}

func (z *ZtreamStorer) StoreZtream(zt *domain.Ztream) error {
	var (
		setActiveZstream = `
		INSERT INTO actives (
			type,
		    reference
			) VALUES (
			    'ztream',
				:name
			)
		ON CONFLICT (type) DO UPDATE SET
			reference = :name
		`
	)
	dbZt := state.NewZtreamDb(zt)
	if _, err := z.db.NamedExec(state.InsertZtream, &dbZt); err != nil {
		return err
	}

	if _, err := z.db.NamedExec(setActiveZstream, &dbZt); err != nil {
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
