package database

import (
	"runtime"

	"github.com/boltdb/bolt"
)

func osPath(file string) string {
	if runtime.GOOS == "windows" {
		return "%appdata%/" + file
	} else {
		return "/home/" + file
	}
}

func Open() (*bolt.DB, error) {
	// Open the botnet.db data file
	// It will be created if it doesn't exist.
	return bolt.Open(osPath("garrison.db"), 0600, nil)
}
