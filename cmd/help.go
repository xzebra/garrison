package cmd

import "fmt"

var (
	helpDesc = "prints this help dialog showing command descriptions"
)

func cmdHelp(args []string) error {
	fmt.Println()
	for name, command := range commandList {
		fmt.Printf("%s - %s\n", name, command.Desc)
	}
	fmt.Println()
	return nil
}
