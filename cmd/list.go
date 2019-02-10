package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"../database"
	"../output"
)

var (
	listDesc = []string{"prints the table of bots stored in the database",
		"Usage: list"}
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
			fmt.Print(output.Center(header, colWidth[i]) + " ")
		}
		fmt.Println()
		for _, w := range colWidth {
			fmt.Print(strings.Repeat("-", w))
			fmt.Print(" ")
		}
		fmt.Println()
		for _, bot := range list {
			fmt.Print(output.Center(strconv.FormatUint(bot.ID, 10), colWidth[0]) + " ")
			fmt.Print(output.Center(bot.Addr, colWidth[1]) + " ")
			fmt.Print(output.Center(bot.Port, colWidth[2]) + " ")

			status := "OFF"
			if bot.Status {
				status = "ON"
			}
			fmt.Print(output.Center(status, colWidth[3]) + "\n")
		}
		fmt.Println()
	}
	return nil
}
