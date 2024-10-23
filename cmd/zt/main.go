package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/adapter/state"
	"github.com/fehlhabers/zt/internal/domain"
)

var Version string = "0.1.0"

func main() {
	log.SetReportTimestamp(false)

	homeDir, _ := os.UserHomeDir()
	dbPath := fmt.Sprintf("%s/.local/state", homeDir)
	storer := state.NewZtreamStorer(dbPath)
	zt := domain.Ztream{Participants: "kaj", Name: "test"}
	storer.StoreZtream(zt)
	gotZt, _ := storer.GetZtream()
	fmt.Println(gotZt)
	// err := cmd.NewRoot().Execute()
	// if err != nil {
	// 	os.Exit(1)
	// }

}
