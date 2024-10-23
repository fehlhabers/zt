package model

const CreateTeamsTable = `	
	CREATE TABLE IF NOT EXISTS teams (
		name text PRIMARY KEY
	)`

type Team struct {
	Name string `db:"name"`
}

const CreateActiveTeamTable = `
	CREATE TABLE IF NOT EXISTS active_team (
		only integer PRIMARY KEY,
	    name text REFERENCES teams
	)`
const CreateActiveZtreamTable = `
	CREATE TABLE IF NOT EXISTS active_ztream (
		only integer PRIMARY KEY,
	    name text REFERENCES ztreams
	)`

const CreateZtreamsTable = `
	CREATE TABLE IF NOT EXISTS ztreams (
		name text PRIMARY KEY,
	    started numeric,
		ends numeric,
	    team text REFERENCES teams
	)`

const InsertZtream = `
		INSERT INTO ztreams (
				name,
				started, 
				ends, 
				team
		) values (
				:name,
				:started,
				:ends,
				:team
		)
		ON CONFLICT (name) DO UPDATE SET
				started=:started,
				ends=:ends,
				team=:team
		`

type Ztream struct {
	Name    string `db:"name"`
	Started int64  `db:"started"`
	Ends    int64  `db:"ends"`
	Team    string `db:"team"`
}

const CreateMembersTable = `
	CREATE TABLE IF NOT EXISTS members (
		name text PRIMARY KEY,
		team text REFERENCES teams
	)`

type Member struct {
	Name string `db:"name"`
	Team string `db:"team"`
}
