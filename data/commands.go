package data

import (
	"errors"
	"fun_stuff/utils"
	"github.com/boltdb/bolt"
	"strings"
)

func addEntry(storePath string) error {
	db, _ := bolt.Open(storePath, 0600, &bolt.Options{Timeout:1000})
	defer db.Close()

	var fields []string
	if err := db.View(func(tx *bolt.Tx) error {
		meta := tx.Bucket([]byte("__meta"))
		if meta == nil {
			return errors.New("uninitialized store")
		}
		rawFields := meta.Get([]byte("fields"))
		fields = strings.Split(string(rawFields), " ")
		return nil
	}); err != nil {
		return err
	}
	name := utils.GetInlineInput("Entry name", "No spaces allowed, cannot be blank", utils.IsProperName)
	return db.Update(func(tx *bolt.Tx) error {
		b, e := tx.CreateBucket([]byte(name))
		if e != nil {
			return e
		}
		for _, field := range fields {
			value := utils.GetInlineInput(field, "Cannot be blank", func(s string) bool {
				return !strings.Contains(s, " ") && s != ""
			})
			if e := b.Put([]byte(field), []byte(value)); e != nil {
				return e
			}
		}
		return nil
	})
}
