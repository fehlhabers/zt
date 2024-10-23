package state

import (
	"fmt"
	"os"
	"testing"

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

func clearDb() {
	dbFile := fmt.Sprintf("%s/ztream.db", dbPath)
	os.Remove(dbFile)
}
