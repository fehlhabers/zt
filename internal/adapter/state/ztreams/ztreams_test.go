package ztreams

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/fehlhabers/zt/internal/domain"
	_ "github.com/glebarez/go-sqlite"
	"github.com/stretchr/testify/assert"
)

var (
	dbPath = os.TempDir()
)

func TestNewZtreamStorer(t *testing.T) {
	clearDb()

	t.Run("create db", func(t *testing.T) {
		got := NewZtreamStorer(dbPath)
		assert.NotNil(t, got)
	})
}

func TestZtream(t *testing.T) {
	clearDb()
	sut := NewZtreamStorer(dbPath)
	t.Run("store new ztream", func(t *testing.T) {

		zt := &domain.Ztream{
			Name:    "zt",
			Started: time.Now(),
			Ends:    time.Now().Add(time.Minute * 15),
		}
		err := sut.StoreZtream(zt)
		assert.NoError(t, err)

		got, err := sut.GetActiveZtream()
		assert.NoError(t, err)
		fmt.Println(got)
		assert.Equal(t, zt.Ends.Unix(), got.Ends.Unix())
	})

	t.Run("update existing team", func(t *testing.T) {

		zt := &domain.Ztream{
			Name:    "zt",
			Started: time.Now(),
			Ends:    time.Now().Add(time.Minute * 15),
		}
		err := sut.StoreZtream(zt)
		assert.NoError(t, err)

		got, err := sut.GetActiveZtream()
		assert.NoError(t, err)
		assert.Equal(t, zt.Ends.Unix(), got.Ends.Unix())
	})
}

func clearDb() {
	dbFile := fmt.Sprintf("%s/ztream.db", dbPath)
	os.Remove(dbFile)
}
