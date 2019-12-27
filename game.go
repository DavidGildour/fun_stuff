package main

import (
	"fmt"
	"fun_stuff/models/items"
	"fun_stuff/models/monsters"
	"fun_stuff/utils"
)

func main() {
	rand := utils.GetRand()

	name := utils.GetAdventurerName()
	prof := utils.GetAdventurerProf()

	player := prof(name)

	fmt.Print(player)
	rarity := items.Rarity[rand.Intn(len(items.Rarity))]
	player.Equip(items.Sword(rarity))
	player.Fight(monsters.RandomMonster())
	fmt.Print(player)
}
