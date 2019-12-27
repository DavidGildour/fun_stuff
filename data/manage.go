package data

import (
	"fmt"
	"fun_stuff/utils"
	"github.com/boltdb/bolt"
)


func executeCommand(storePath, cmdWithArgs string) error {
	cmd, _ := parseCommand(cmdWithArgs)
	switch cmd {
	case "add":
		return addEntry(storePath)
	default:
		return nil
	}
}

func Manage(store string) {
	fmt.Printf("Connecting to the \"%s\" store...\n", store)
	storePath := storePath(store)
	db, err := bolt.Open(storePath, 0600, &bolt.Options{Timeout:1000})
	if err != nil {
		fmt.Printf("Error opening a store.\nError message: %s\nShutting down...\n", err)
		return
	}

	choice := utils.GetUserInput("Connection OK. Print out list of entries? (Y/n)", "", nil)
	switch choice {
	case "Y", "y", "":
		if err = printEntries(db); err != nil {
			fmt.Printf("Error reading a store.\nError message: %s\nShutting down...\n", err)
			return
		}
	}
	if err = db.Close(); err != nil {
		return
	}
	for {
		cmd := utils.GetUserInput("", "Unknown command", isProperCommand)
		if cmd == "q" {
			return
		}
		if err := executeCommand(storePath, cmd); err != nil {
			panic(err)
		}
	}
}
