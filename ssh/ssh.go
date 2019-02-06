package ssh

import "strconv"

func CheckConnection(ip string, port string) (bool, error) {
	return true, nil
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
