package cmd

import (
	"fmt"
	"strings"
	"sync"

	"../database"
	"../output"
	"../ssh"
)

var (
	beaconDesc = []string{"broadcast a command to all the online bots",
		"Usage: beacon {command}"}
)

func cmdBeacon(args []string) error {
	bots, err := database.ListBots()
	if err != nil {
		return err
	}

	command := strings.Join(args, " ")
	count := 0
	c := make(chan bool)
	var wg sync.WaitGroup
	for _, bot := range bots {
		if bot.Status == true {
			wg.Add(1)
			count++
			go ssh.SendCommand(database.Listed2Bot(&bot), command, &wg, c)
		}
	}

	successfulBots := 0
	for i := 0; i < count; i++ {
		success := <-c
		if success {
			successfulBots++
		}
	}

	wg.Wait()
	output.Info(fmt.Sprintf("%d/%d bots completed the task", successfulBots, count))
	return nil
}
