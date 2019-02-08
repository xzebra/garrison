package cmd

import "fmt"

type cmdFunc func([]string) error

type Cmd struct {
	Desc string
	Func cmdFunc
}

var (
	commandList map[string]Cmd
	UnknownCmd  = fmt.Errorf("Unknown command")
	WrongArgs   = fmt.Errorf("Wrong argument value or syntax")
	UnknownArgs = fmt.Errorf("Unknown argument")
)

func Init() {
	commandList = map[string]Cmd{
		"help":  Cmd{helpDesc, cmdHelp},
		"quit":  Cmd{quitDesc, cmdQuit},
		"add":   Cmd{addDesc, cmdAdd},
		"list":  Cmd{listDesc, cmdList},
		"del":   Cmd{delDesc, cmdDel},
		"shell": Cmd{shellDesc, cmdShell},
	}
}

func Execute(name string, args []string) error {
	if command, found := commandList[name]; found {
		return command.Func(args)
	}
	return UnknownCmd
}

func parseArgs(args []string) map[string]string {
	output := make(map[string]string)
	for i := 0; i < len(args)-1; i++ {
		if args[i][0] == '-' {
			output[args[i][1:]] = args[i+1]
			i++
		}
	}
	return output
}
