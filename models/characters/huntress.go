package characters

import (
	"fun_stuff/models/items"
	"strings"
)

func Huntress(name string) Character {
	return &genericCharacter{
		_adj:      "swift",
		_type:     "huntress",
		_deity:    "blessed Angaram",
		Name:      strings.Title(name),
		Hp:        80,
		BaseDmg:   7,
		BaseArmor: 0,
		Equipment: []items.Item{},
	}
}
