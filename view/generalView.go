package view

import (
	"bufio"
	"fmt"
	"fun_stuff/utils"
	"os"
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
	for {
		fmt.Print(">> ")
		scanner.Scan()
		input := scanner.Text()
		if utils.IsProperChoice(input, len(choices)) {
			return input
		}
		fmt.Println("Wrong number, try again.")
	}
}
