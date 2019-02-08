package database

import (
	"encoding/json"

	"github.com/boltdb/bolt"
)

type Bot struct {
	Addr   string
	User   string
	Pwd    string
	Port   string
	Status bool
}

type ListedBot struct {
	ID     uint64
	Addr   string
	User   string
	Pwd    string
	Port   string
	Status bool
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
			list = append(list, ListedBot{})
			json.Unmarshal(v, &list[i])
			list[i].ID = btoi(k)
			i++
		}

		return nil
	})
}

func GetBot(id uint64) (Bot, error) {
	var bot Bot
	var botJSON []byte

	err := Session.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		botJSON = b.Get(itob(id))
		return nil
	})
	if err != nil {
		return bot, err
	}

	err = json.Unmarshal(botJSON, &bot)
	return bot, err
}
