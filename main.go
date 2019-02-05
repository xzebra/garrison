package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./cmd"
	"./utils"
)

func main() {
	printTitle()
	cmd.Init()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		utils.OutputCursor()
		if !scanner.Scan() {
			utils.OutputError(scanner.Err())
			return
		}

		input := strings.Split(scanner.Text(), " ")
		err := cmd.Execute(input[0], input[1:])
		if err != nil {
			utils.OutputError(err)
		}

		fmt.Println()
	}
}
