package ssh

import (
	"strconv"

	"../database"
	"golang.org/x/crypto/ssh"
)

var (
	ErrConn = "could not connect to host"
)

func createConn(bot *database.Bot) (*ssh.Client, error) {
	var hostKey ssh.PublicKey
	config := &ssh.ClientConfig{
		User: bot.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(bot.Pwd),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	return ssh.Dial("tcp", bot.Addr+":"+bot.Port, config)
}

func CheckConnection(bot *database.Bot) (bool, error) {
	conn, err := createConn(bot)
	if err != nil || conn == nil {
		return false, err
	}
	defer conn.Close()
	return true, nil
}

func CreateShell(bot *database.Bot) error {
	conn, err := createConn(bot)
	if err != nil {
		return err
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	// request terminal
	if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
		return err
	}
	// start remote shell
	if err := session.Shell(); err != nil {
		return err
	}
	return nil
}

func IsValidAddr(addr string) bool {
	return true
}

func IsValidPort(port string) bool {
	if len(port) == 0 {
		return false
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		return false
	}
	return p > 0 && p < 65535
}
