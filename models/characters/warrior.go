package characters

import (
	"fun_stuff/models/items"
	"strings"
)

func Warrior(name string) Character {
	return &genericCharacter{
		_adj:      "mighty",
		_type:     "warrior",
		_deity:    "almighty Volcatosh",
		name:      strings.Title(name),
		hp:        100,
		baseDmg:   5,
		baseArmor: 1,
		equipment: []items.Item{},
	}
}
