package state

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/fehlhabers/zt/internal/model"
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

		zt := model.Ztream{
			Name:    "zt",
			Started: time.Now().Unix(),
			Ends:    time.Now().Unix() + 60*15,
			Team:    "test1",
		}
		err := sut.StoreZtream(zt)
		assert.NoError(t, err)

		got, err := sut.GetActiveZtream()
		assert.NoError(t, err)
		fmt.Println(got)
		assert.Equal(t, zt.Ends, got.Ends)
	})

	t.Run("update existing team", func(t *testing.T) {

		zt := model.Ztream{
			Name:    "zt",
			Started: time.Now().Unix(),
			Ends:    time.Now().Unix() + 60*15,
			Team:    "test1",
		}
		err := sut.StoreZtream(zt)
		assert.NoError(t, err)

		got, err := sut.GetActiveZtream()
		assert.NoError(t, err)
		assert.Equal(t, zt.Ends, got.Ends)
	})
}

func clearDb() {
	dbFile := fmt.Sprintf("%s/ztream.db", dbPath)
	os.Remove(dbFile)
}
