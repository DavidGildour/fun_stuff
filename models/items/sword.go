package items


func Sword(rarity string) Item {
	var dmg int64
	switch rarity {
	case "broken":
		dmg = 3
	case "rare":
		dmg = 7
	case "epic":
		dmg = 10
	case "legendary":
		dmg = 15
	default:
		dmg = 5
	}
	return Item{
		Rarity: rarity,
		Name: "Sword",
		Dmg:  dmg,
	}
}
