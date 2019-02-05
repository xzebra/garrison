package database

import (
	"log"
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

func Open() {
	// Open the botnet.db data file
	// It will be created if it doesn't exist.
	db, err := bolt.Open(osPath("my.db"), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
