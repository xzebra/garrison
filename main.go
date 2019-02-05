package main

import (
	"bufio"
	"fmt"
	"os"

	"./cmd"
)

func main() {
	printTitle()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(":: ")
		if !scanner.Scan() {
			fmt.Printf("[!] %s\n", scanner.Err())
			return
		}

		switch scanner.Text() {
		case "quit":
			return
		case "help":
			cmd.Help()
		}

		fmt.Println()
	}
}
