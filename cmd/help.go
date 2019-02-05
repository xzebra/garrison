package cmd

import "fmt"

func Help() {
	for name, desc := range commandList {
		fmt.Printf("%s - %s\n", name, desc)
	}
}
