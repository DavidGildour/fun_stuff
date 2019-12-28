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
		name:      strings.Title(name),
		hp:        80,
		baseDmg:   7,
		baseArmor: 0,
		equipment: []items.Item{},
	}
}
