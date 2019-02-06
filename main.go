package main

import (
	"bufio"
	"os"
	"strings"

	"./cmd"
	"./database"
	"./output"
)

func main() {
	printTitle()

	cmd.Init()
	err := database.Init()
	if err != nil {
		output.Error(err)
		return
	}
	defer database.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		output.Cursor()
		if !scanner.Scan() {
			output.Error(scanner.Err())
			return
		}

		input := strings.Split(scanner.Text(), " ")
		err := cmd.Execute(input[0], input[1:])
		if err != nil {
			output.Error(err)
		}
	}
}
