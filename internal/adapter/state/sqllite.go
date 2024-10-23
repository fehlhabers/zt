package state

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/domain"
	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
)

var _ domain.Storer = &ZtreamStorer{}

type ZtreamStorer struct {
	db *sqlx.DB
}

const createZtreamsTable = `
	CREATE TABLE IF NOT EXISTS ztreams (
		name text,
	    id text, 
		active boolean,
	    team text
	)`

type Ztream struct {
	name   string `db:"name"`
	id     string `db:"id"`
	active bool   `db:"active"`
	team   string `db:"team"`
}

const createMembersTable = `	
	CREATE TABLE IF NOT EXISTS members (
		name text,
		team text
	)`

type Member struct {
	name string `db:"name"`
	team string `db:"team"`
}

const createTeamsTable = `	
	CREATE TABLE IF NOT EXISTS teams (
		id text
		default_team boolean
	)`

type Team struct {
	id           string `db:"id"`
	default_team bool   `db:"default_team"`
}

func NewZtreamStorer(dbPath string) *ZtreamStorer {

	mustCreateDirectories(dbPath)

	dbFile := fmt.Sprintf("%s/ztream.db", dbPath)
	db, err := sqlx.Open("sqlite", dbFile)
	if err != nil {
		log.Fatal("Not able to create db instance", "path", dbPath, "error", err)
	}

	if _, err := db.Exec(createZtreamsTable); err != nil {
		log.Fatal("Not able to prepare ztream", "error", err)
	}

	if _, err := db.Exec(createMembersTable); err != nil {
		log.Fatal("Not able to prepare members", "error", err)
	}

	if _, err := db.Exec(createTeamsTable); err != nil {
		log.Fatal("Not able to prepare team", "error", err)
	}
	return &ZtreamStorer{db: db}
}

// GetZtream implements domain.Storer.
func (z *ZtreamStorer) GetZtream() (domain.Ztream, error) {
	row := z.db.QueryRow("select participants, name from ztreams")
	var zt domain.Ztream
	row.Scan(&zt.Participants, &zt.Name)

	return zt, nil
}

// StoreZtream implements domain.Storer.
func (z *ZtreamStorer) StoreZtream(zt domain.Ztream) error {

	_, err := z.db.Exec("insert into ztreams ('participants', 'name') values (?,?)", zt.Participants, zt.Name)
	return err
}

func mustCreateDirectories(path string) {
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		log.Fatal("Not able to create directories to store ztream data", "path", path, "error", err)
	}
}
