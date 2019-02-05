package cmd

import "fmt"

var (
	helpDesc = "prints this help dialog showing command descriptions"
)

func cmdHelp(args []string) error {
	for name, command := range commandList {
		fmt.Printf("%s - %s\n", name, command.Desc)
	}
	return nil
}
