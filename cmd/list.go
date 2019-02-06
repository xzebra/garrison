package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"../database"
	"../output"
)

var (
	listDesc    = "prints the table of bots stored in the database"
	listHeaders = []string{"ID", "Address", "Port", "Status"}
	colWidth    = []int{5, 16, 6, 6}
)

func cmdList(args []string) error {
	list, err := database.ListBots()
	if err != nil {
		return err
	}
	if len(list) == 0 {
		output.Info("no bots in database")
	} else {
		fmt.Println()
		for i, header := range listHeaders {
			fmt.Print(center(header, colWidth[i]) + " ")
		}
		fmt.Println()
		for _, w := range colWidth {
			fmt.Print(strings.Repeat("-", w))
			fmt.Print(" ")
		}
		fmt.Println()
		for _, bot := range list {
			fmt.Print(center(strconv.FormatUint(bot.ID, 10), colWidth[0]) + " ")
			fmt.Print(center(bot.Addr, colWidth[1]) + " ")
			fmt.Print(center(bot.Port, colWidth[2]) + " ")

			status := "OFF"
			if bot.Status {
				status = "ON"
			}
			fmt.Print(center(status, colWidth[3]) + "\n")
		}
		fmt.Println()
	}
	return nil
}

func center(input string, width int) string {
	var output string
	dif := width - len(input)
	side := dif / 2
	if dif%2 != 0 {
		output += " "
	}
	output += strings.Repeat(" ", side)
	output += input
	output += strings.Repeat(" ", side)
	return output
}
