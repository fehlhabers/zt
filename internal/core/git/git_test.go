package git

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPull(t *testing.T) {
	stdout, err := Pull()
	fmt.Println(stdout)
	assert.NoError(t, err)
}

func TestAddAll(t *testing.T) {
	stdout, err := AddAll()
	fmt.Println(stdout)
	assert.NoError(t, err)
}

func TestCommit(t *testing.T) {
	stdout, err := Commit("WiP")
	fmt.Println(stdout)
	assert.NoError(t, err)
}

func TestCurrentBranch(t *testing.T) {
	stdout, err := CurrentBranch()
	fmt.Println(stdout)
	assert.NoError(t, err)
}

func TestPush(t *testing.T) {
	stdout, err := Push()
	fmt.Println(stdout)
	assert.NoError(t, err)
}

func TestCreateBranch(t *testing.T) {
	stdout, err := CreateBranch("test-branch")
	fmt.Println(stdout)
	assert.NoError(t, err)
}
