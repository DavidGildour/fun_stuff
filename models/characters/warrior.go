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
		Name:      strings.Title(name),
		Hp:        100,
		BaseDmg:   5,
		BaseArmor: 1,
		Equipment: []items.Item{},
	}
}
