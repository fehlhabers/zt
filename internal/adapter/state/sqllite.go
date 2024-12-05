package state

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
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

	dbFile := fmt.Sprintf("%s/state.db", dbPath)
	db, err := sqlx.Open("sqlite", dbFile)
	if err != nil {
		log.Fatal("Not able to create db instance", "path", dbPath, "error", err)
	}

	if _, err := db.Exec(CreateZtreamsTable); err != nil {
		log.Fatal("Not able to prepare ztream", "error", err)
	}

	if _, err := db.Exec(CreateActivesTable); err != nil {
		log.Fatal("Not able to prepare actives", "error", err)
	}

	return &ZtreamStorer{db: db}
}

func (z *ZtreamStorer) GetAllZtreams() []Ztream {
	var (
		selectAllZtreams = `
		SELECT z.* FROM ztreams z
		`
	)
	rows, err := z.db.Query(selectAllZtreams)
	if err != nil {
		return nil
	}

	var ztreams []Ztream
	for rows.Next() {
		var z Ztream

		err := rows.Scan(&z.Name, &z.Started, &z.Ends, &z.Team)
		if err != nil {
			log.Fatal(err)
		}

		ztreams = append(ztreams, z)
	}

	return ztreams
}

func (z *ZtreamStorer) GetActiveZtream() (Ztream, error) {
	var zt Ztream
	var (
		selectActiveZtream = `
		SELECT z.* FROM ztreams z INNER JOIN actives a ON (a.reference = z.name) WHERE a.type = 'ztream'
		`
	)

	if err := z.db.Get(&zt, selectActiveZtream); err != nil {
		return zt, err
	}
	return zt, nil
}

func (z *ZtreamStorer) StoreZtream(zt Ztream) error {
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
	if _, err := z.db.NamedExec(InsertZtream, &zt); err != nil {
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
