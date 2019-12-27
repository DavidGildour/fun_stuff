package monsters

func Skeleton() *Monster {
	return &Monster{
		Name: "skeleton",
		Dmg: 7,
		Hp: 10,
		Armor: 1,
	}
}
