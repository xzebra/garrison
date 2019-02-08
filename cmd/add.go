package cmd

import (
	"fmt"

	"../database"
	"../output"
	"../ssh"
)

var (
	addDesc        = "adds a bot to the database\n\t[-c] only adds it if check login is successful"
	ErrAddrInvalid = fmt.Errorf("address not valid")
	ErrPortInvalid = fmt.Errorf("port not valid")
	ErrNoConn      = fmt.Errorf("couldn't connect to the host")
)

func cmdAdd(args []string) error {
	if len(args) == 0 {
		return ErrNoArgs
	}
	flags := parseArgs(args)
	if !ssh.IsValidAddr(args[0]) {
		return ErrAddrInvalid
	}

	// parse port flag or assign default 22 port
	port, found := flags["p"]
	if found {
		if !ssh.IsValidPort(port) {
			return ErrPortInvalid
		}
	} else {
		port = "22"
	}

	bot := database.Bot{
		Addr:   args[0],
		User:   flags["user"],
		Pwd:    flags["pwd"],
		Port:   port,
		Status: false,
	}

	status, err := ssh.CheckConnection(&bot)
	if flags["c"] == "true" && err != nil {
		// only add bot if check connection was successful
		return ErrNoConn
	}
	bot.Status = status

	if err := database.AddBot(&bot); err != nil {
		return err
	}

	output.Added("added bot to database")
	return nil
}
