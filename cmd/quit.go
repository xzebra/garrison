package cmd

import "os"

var (
	quitDesc = "exits the program"
)

func cmdQuit(args []string) error {
	os.Exit(0)
	return nil
}
