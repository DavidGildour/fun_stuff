package data

import (
	"fmt"
	"github.com/boltdb/bolt"
	"strconv"
	"strings"
)

const (
	MAXWIDTH int = 20
	ITEMSPERLINE int = 4
)

func storePath(s string) string {
	return STOREDIR + s + ".db"
}

func printEntries(db *bolt.DB) error {
	var items []string
	err := db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			if !strings.Contains(string(name), "__") {
				items = append(items, string(name))
			}
			return nil
		})
	})
	if err != nil {
		return err
	}
	for i, name := range items {
		if i > 0 && i % ITEMSPERLINE == 0 {
			fmt.Println()
		}
		indexedName := strconv.FormatInt(int64(i)+1, 10) + ". " + name
		padding := MAXWIDTH - len(indexedName)
		fmt.Print(indexedName + strings.Repeat(" ", padding))
	}
	fmt.Println()
	return nil
}

func isProperCommand(s string) bool {
	cmd := strings.Fields(s)[0]
	switch strings.ToLower(cmd) {
	case "add", "edit", "remove", "rm", "q", "ls", "list", "show":
		return true
	default:
		return false
	}
}

func parseCommand(cmd string) (string, []string) {
	all := strings.Fields(cmd)
	return strings.ToLower(all[0]), all[1:]
}