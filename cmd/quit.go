package cmd

import "os"

var (
	quitDesc = []string{"exits the program",
		"Usage: quit"}
)

func cmdQuit(args []string) error {
	os.Exit(0)
	return nil
}
