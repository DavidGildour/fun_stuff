package main

import (
	"fmt"
	"fun_stuff/data"
	"fun_stuff/utils"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		store := utils.GetUserInput(
			"Choose a store to manage: ",
			"No such store.",
			utils.IsProperStore)
		data.Manage(store)
	} else {
		if args[0] == "init" {
			fmt.Println("Initializing stores...\n")
			if err := data.InitStores(); err != nil {
				panic(err)
			}
			fmt.Println("\nStores initialized successfully.")
			return
		}
		store := args[0]
		if utils.IsProperStore(store) {
			data.Manage(args[0])
		} else {
			fmt.Println("No such store.")
		}
	}
}
