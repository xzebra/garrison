package database

import (
	"encoding/binary"
	"encoding/json"
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

func AddBot(bot *Bot) error {
	return Session.Update(func(tx *bolt.Tx) error {
		// Retrieve the bots bucket
		b := tx.Bucket(bucketName)

		// Generate ID for the bot.
		// This returns an error only if the Tx is closed or not writeable
		// That can't happen in an Update() call so ignore the error check
		id, _ := b.NextSequence()

		// Marshal bot data into bytes
		buf, err := json.Marshal(bot)
		if err != nil {
			return err
		}

		// Persist bytes to bots bucket
		return b.Put(itob(id), buf)
	})
}

func RemoveBot(id uint64) error {
	return Session.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		// id integer representation to database key
		key := itob(id)
		return b.Delete(key)
	})
}

func ListBots() ([]ListedBot, error) {
	var list []ListedBot
	return list, Session.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket(bucketName)
		c := b.Cursor()

		i := 0
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var bot Bot
			json.Unmarshal(v, &bot)
			list = append(list, ListedBot{
				ID:     btoi(k),
				Addr:   bot.Addr,
				Pwd:    bot.Pwd,
				Port:   bot.Port,
				Status: bot.Status,
			})
			i++
		}

		return nil
	})
}

func Close() {
	Session.Close()
}
