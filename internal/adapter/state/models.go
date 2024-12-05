package state

const CreateActivesTable = `
	CREATE TABLE IF NOT EXISTS actives (
		type text PRIMARY KEY,
	    reference text REFERENCES ztreams
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