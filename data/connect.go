package data

import (
	"fmt"
	"strings"

	"github.com/boltdb/bolt"
)

const STOREDIR = "./data/stores/"

var (
	places = StoreInfo{
		Name:   "places",
		Path:   STOREDIR + "places.db",
		Fields: []string{"name", "monsters"},
	}
	items = StoreInfo{
		Name:   "items",
		Path:   STOREDIR + "items.db",
		Fields: []string{"rarity", "name", "dmg", "armor"},
	}
	monsters = StoreInfo{
		Name:   "monsters",
		Path:   STOREDIR + "monsters.db",
		Fields: []string{"name", "dmg", "armor", "hp"},
	}
	characters = StoreInfo{
		Name:   "characters",
		Path:   STOREDIR + "characters.db",
		Fields: []string{"adj", "name", "base_dmg", "base_armor", "deity", "base_hp"},
	}
)

func initStore(store StoreInfo) error {
	db, err := bolt.Open(store.Path, 0600, &bolt.Options{Timeout: 1000})
	if err != nil {
		return err
	}
	fmt.Print(" Store file created...")
	if err = db.Update(func(tx *bolt.Tx) error {
		b, e := tx.CreateBucketIfNotExists([]byte("__meta"))
		if e != nil {
			return e
		}
		e = b.Put([]byte("fields"), []byte(strings.Join(store.Fields, " ")))
		if e != nil {
			return e
		}
		fmt.Print(" Meta information written...")
		return nil
	}); err != nil {
		return err
	}
	if err = db.Close(); err != nil {
		return err
	}
	return nil
}

func InitStores() error {
	stores := []StoreInfo{places, items, characters, monsters}
	for i, store := range stores {
		fmt.Printf("%d/%d %s...", i+1, len(stores), store.Name)
		if err := initStore(store); err != nil {
			fmt.Printf("Error initializing \"%s\"! Aborting.", store.Name)
			return err
		}
		fmt.Println(" Done.")
	}
	return nil
}
