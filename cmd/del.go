package cmd

import (
	"fmt"
	"strconv"

	"../database"
	"../output"
)

var (
	delDesc   = "removes bot data from database given its id"
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
