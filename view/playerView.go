package view

import (
	"fmt"
	"fun_stuff/models/characters"
	"fun_stuff/models/items"
)

// EquipInfo is a method for displaying info about item equipped by character (cName)
func EquipInfo(cName string, i items.Item) {
	fmt.Printf("%s equips a %s %s!\n", cName, i.Rarity, i.Name)
}

// PrintPlayerInfo is a function for printing plaer info :p
func PrintPlayerInfo(c characters.Character) {
	str := fmt.Sprintf("A %s %s - %s!", c.Adj(), c.Type(), c.Name())
	if c.Hp() > 0 {
		str += fmt.Sprintf(" As of now, he has %d HP, %d armor and deals %d damage per hit.\n", c.Hp(), c.Armor(), c.Dmg())
		if len(c.Equipment()) > 0 {
			str += "Equipment:\n"
			for _, item := range c.Equipment() {
				str += " > " + fmt.Sprintln(item)
			}
		}
	} else {
		str += " Apparently dead."
	}
	fmt.Print(str)
}
