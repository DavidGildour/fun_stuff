package monsters

type (
	Monster struct {
		Name string
		Hp int64
		Armor int64
		Dmg int64
	}
)

func RandomMonster() *Monster {
	for _, m := range Monsters {
		return &m
	}
	return nil
}

var Monsters = []Monster{
	{
		Name:  "goblin",
		Hp:    15,
		Armor: 0,
		Dmg:   3,
	},
	{
		Name: "skeleton",
		Hp: 12,
		Armor: 1,
		Dmg: 4,
	},
}

