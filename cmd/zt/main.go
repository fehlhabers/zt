package main

import (
	"os"

	"github.com/fehlhabers/st/internal/cmd"
)

func main() {
	err := cmd.NewRoot().Execute()
	if err != nil {
		os.Exit(1)
	}

}
