package database

import (
	"encoding/binary"
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/boltdb/bolt"
)

var (
	Path       = osPath("garrison.db")
	Session    *bolt.DB
	bucketName = []byte("bots")
)

func osPath(file string) string {
	if runtime.GOOS == "windows" {
		return path.Join(os.Getenv("userprofile"), "Documents", file)
	}
	return path.Join("/home/", file)
}

// itob returns an 8-byte big endian representation of v
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// btoi returns an integer representation of v
func btoi(v []byte) uint64 {
	return binary.BigEndian.Uint64(v)
}

func createBucket(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		var err error
		_, err = tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return fmt.Errorf("Couldn't create table: %s", err)
		}
		return nil
	})
}

func Init() error {
	var err error
	Session, err = bolt.Open(Path, 0644, nil)
	if err != nil {
		return err
	}

	err = createBucket(Session)
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	Session.Close()
}
