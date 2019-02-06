package cmd

import (
	"fmt"

	"../database"
	"../output"
	"../ssh"
)

var (
	addDesc        = "adds a bot to the database\n\t[-c] only adds it if check login is successful"
	ErrAddrInvalid = fmt.Errorf("Address not valid")
)

func cmdAdd(args []string) error {
	if !ssh.IsValidAddr(args[0]) {
		return ErrAddrInvalid
	}
	conn, err := ssh.CheckConnection(args[0], args[1])
	if err != nil {
		return err
	}

	bot := database.Bot{
		Addr:   args[0],
		Pwd:    args[1],
		Port:   args[2],
		Status: conn,
	}
	if err = database.AddBot(&bot); err != nil {
		return err
	}

	output.Added("added bot to database")
	return nil
}
