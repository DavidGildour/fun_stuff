package utils

import (
	"bufio"
	"fmt"
	"fun_stuff/models/characters"
	"os"
	"regexp"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func IsProperStore(s string) bool {
	switch strings.ToLower(s) {
	case "items", "monsters", "places", "characters":
		return true
	default:
		return false
	}
}

func IsProperName(s string) bool {
	return regexp.MustCompile(`^[[:alpha:]]+$`).MatchString(s)
}

func isProperProf(s string) bool {
	_, ok := characters.Professions[s]
	return ok
}

func GetUserInput(prompt, err string, validator func(string) bool) string {
	if prompt != "" {
		fmt.Println(prompt)
	}
	for {
		fmt.Print(">> ")
		scanner.Scan()
		input := scanner.Text()
		if validator == nil || validator(input) {
			return input
		}
		fmt.Printf("%s: %s\n", err, input)
	}
}

func GetInlineInput(prompt, err string, validator func(string) bool) string {
	for {
		fmt.Print(strings.Title(prompt) + ": ")
		scanner.Scan()
		input := scanner.Text()
		if validator(input) {
			return input
		}
		fmt.Print(err + "\n")
	}
}

func GetAdventurerName() string {
	return GetUserInput("What shall be thy name?", "That's not a name...", IsProperName)
}

func GetAdventurerProf() characters.Profession {
	choice := GetUserInput(
		"To which order do you belong?\n1. Warrior 2. Huntress",
		"That's not an option",
		isProperProf)
	return characters.Professions[choice]
}
