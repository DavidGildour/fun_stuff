package view

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const HEADER string = "\n====== +++++++ =======\n"
const INDENT string = "     "

var scanner = bufio.NewScanner(os.Stdin)

func PrettyPrint(line string, br bool, head bool) {
	if br == true {
		line += "\n"
	}
	if head == true {
		fmt.Println(INDENT + HEADER)
	}
	fmt.Printf("%s%s", INDENT, line)
}

func GetPlayerActionChoice(choices []string) string {
	PrettyPrint("Choose an action:", true, true)
	for i, choice := range choices {
		PrettyPrint(fmt.Sprintf("%s%d.) %s", INDENT, i+1, choice), true, false)
	}
	var validator = func(s string) bool {
		if !regexp.MustCompile(`^\d+$`).MatchString(s) {
			return false
		}
		num, _ := strconv.Atoi(s)
		if num > len(choices) || num < 1 {
			return false
		}
		return true
	}
	for {
		fmt.Print(">> ")
		scanner.Scan()
		input := scanner.Text()
		if validator(input) {
			return input
		}
		fmt.Println("Wrong number, try again.")
	}
}
