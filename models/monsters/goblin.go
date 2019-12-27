package monsters

func Goblin() *Monster {
	return &Monster{
		Name:  "goblin",
		Hp:    15,
		Armor: 0,
		Dmg:   3,
	}
}