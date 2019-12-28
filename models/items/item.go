package items

import (
	"fmt"
	"strings"
)

var Rarity = [5]string{
	"broken", "common", "rare", "epic", "legendary",
}

type Item struct {
	Rarity string
	Name   string
	Dmg    int64
	Armor  int64
}

func (i Item) String() string {
	return fmt.Sprintf("%s %s - %d dmg, %d armor\n", strings.Title(i.Rarity), i.Name, i.Dmg, i.Armor)
}
