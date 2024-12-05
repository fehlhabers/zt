package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Stash() (stdout string, err error) {
	return git("stash")
}

func StashPop() (stdout string, err error) {
	return git("stash", "pop")
}

func SwitchBranch(branch string) (stdout string, err error) {
	return git("checkout", branch)
}

func CreateBranch(branch string) (string, error) {
	return git("checkout", "-b", branch)
}

func CurrentBranch() (string, error) {
	branch, err := git("branch", "--show-current")
	return strings.Trim(branch, "\n"), err
}

func Fetch() (stdout string, err error) {
	return git("fetch", "--all")
}

func Pull() (stdout string, err error) {
	return git("pull")
}

func AddAll() (stdout string, err error) {
	return git("add", "--all")
}

func Commit(commitMessage string) (stdout string, err error) {
	return git("commit", "-m", commitMessage)
}

func Push() (stdout string, err error) {
	return git("push")
}

func PushSetOrigin(branch string) (stdout string, err error) {
	return git("push", "--set-upstream", "origin", branch)
}

func git(arguments ...string) (stdout string, err error) {
	git := make([]string, 1)
	git[0] = "git"
	arguments = append(git, arguments...)
	return runCommand(arguments...)
}

func runCommand(arguments ...string) (stdout string, err error) {
	var stderr, stdoutBuf bytes.Buffer
	cmd := exec.Command(arguments[0], arguments[1:]...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdoutBuf
	err = cmd.Run()
	if err != nil {
		fmt.Println(stderr.String())
		return
	}

	stdout = stdoutBuf.String()
	return
}
