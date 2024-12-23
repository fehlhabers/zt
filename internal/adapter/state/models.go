package state

import (
	"time"

	"github.com/fehlhabers/zt/internal/domain"
)

const CreateActivesTable = `
	CREATE TABLE IF NOT EXISTS actives (
		type text PRIMARY KEY,
	    reference text REFERENCES ztreams
	)`

const CreateZtreamsTable = `
	CREATE TABLE IF NOT EXISTS ztreams (
		name text PRIMARY KEY,
		metadata text,
	    started numeric,
		ends numeric
	)`

const InsertZtream = `
		INSERT INTO ztreams (
				name,
				metadata,
				started, 
				ends
		) values (
				:name,
				:metadata,
				:started,
				:ends
		)
		ON CONFLICT (name) DO UPDATE SET
				started=:started,
				ends=:ends
		`

type ZtreamDb struct {
	Name     string `db:"name"`
	Metadata string `db:"metadata"`
	Started  int64  `db:"started"`
	Ends     int64  `db:"ends"`
}

func (ztream *ZtreamDb) ToZtream() *domain.Ztream {
	return &domain.Ztream{
		Name:     ztream.Name,
		Metadata: ztream.Metadata,
		Started:  time.Unix(ztream.Started, 0),
		Ends:     time.Unix(ztream.Ends, 0),
	}
}

func NewZtreamDb(ztream *domain.Ztream) *ZtreamDb {
	return &ZtreamDb{
		Name:    ztream.Name,
		Started: ztream.Started.Unix(),
		Ends:    ztream.Ends.Unix(),
	}
}
