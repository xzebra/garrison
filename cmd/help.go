package cmd

import (
	"fmt"

	"../output"
)

var (
	helpDesc = []string{"prints this help dialog showing command descriptions",
		"Usage: help [cmd]\n" +
			"[cmd] will print an extended usage description"}
	ErrCmdNotFound = fmt.Errorf("command not found")
)

func cmdHelp(args []string) error {
	if len(args) == 1 {
		return helpCmd(args[0])
	}

	fmt.Println()
	for name, command := range commandList {
		fmt.Printf("%s - %s\n", output.CenterRight(name, MaxCmdWidth), command.Desc[0])
	}
	fmt.Println()
	return nil
}

func helpCmd(name string) error {
	if command, found := commandList[name]; found {
		fmt.Println()
		fmt.Println(command.Desc[1])
		fmt.Println()
		return nil
	}
	return ErrCmdNotFound
}
