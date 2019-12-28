package view

import (
	"fmt"
	"fun_stuff/models/characters"
)

// AnnounceBattle is ...
func AnnounceBattle(cName string, mName string) {
	PrettyPrint(fmt.Sprintf("%s is fighting a filthy %s! The battle unravels...\n", cName, mName), true, true)
}

// AnnounceDefeat ...
func AnnounceDefeat(c characters.Character, mName string) {
	PrettyPrint(fmt.Sprintf("By the gods! A %s %s has been slain by the filthy %s!\n", c.Adj(), c.Name(), mName), true, true)
}

// AnnounceVictory ...
func AnnounceVictory(c characters.Character, mName string) {
	PrettyPrint(fmt.Sprintf("Praise the %s! A filthy %s perishes!\n", c.Deity(), mName), true, true)
}
