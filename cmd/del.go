package cmd

import (
	"fmt"
	"strconv"

	"../database"
	"../output"
)

var (
	delDesc = []string{"removes bot data from database given its id",
		"Usage: del [id] [-c true]\n" +
			"[-c] delete all offline bots\n" +
			"{id} [-c] delete if offline"}
	ErrNoArgs = fmt.Errorf("insuficient arguments")
	ErrSyntax = fmt.Errorf("argument empty or syntax error")
)

func cmdDel(args []string) error {
	if len(args) == 0 {
		return ErrNoArgs
	}
	if len(args[0]) == 0 {
		return ErrSyntax
	}

	flags := parseArgs(args)
	if flags["off"] == "true" {
		if !isFlag(args[0]) {
			return deleteIfOff(args[0])
		}
		return deleteAllOff()
	}

	id, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return err
	}

	err = database.RemoveBot(id)
	if err != nil {
		return err
	}
	output.Removed(fmt.Sprintf("removed bot with id %d", id))
	return nil
}

func deleteAllOff() error {
	bots, err := database.ListBots()
	if err != nil {
		return err
	}
	for _, bot := range bots {
		if bot.Status == false {
			database.RemoveBot(bot.ID)
		}
	}
	output.Removed("removed all offline bots")
	return nil
}

func deleteIfOff(idArg string) error {
	id, err := strconv.ParseUint(idArg, 10, 64)
	if err != nil {
		return err
	}

	bot, err := database.GetBot(id)
	if err != nil {
		return err
	}

	if bot.Status == false {
		database.RemoveBot(id)
	}
	output.Removed(fmt.Sprintf("removed bot with id %d", id))
	return nil
}
