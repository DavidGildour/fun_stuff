package controllers

import (
	"fun_stuff/models/characters"
	"fun_stuff/models/items"
	"fun_stuff/models/monsters"
	"fun_stuff/utils"
	"fun_stuff/view"
)

var player characters.Character = nil

func SetupGame() {
	name := utils.GetAdventurerName()
	prof := utils.GetAdventurerProf()

	player = prof(name)
	rand := utils.GetRand()
	rarity := items.Rarity[rand.Intn(len(items.Rarity))]
	PlayerEquipItem(items.Sword(rarity))
	view.PrintPlayerInfo(player)
}

func PlayerEquipItem(i items.Item) {
	player.Equip(i)
	view.EquipInfo(player.Name(), i)
}

func FightRandomMonster() {
	Fight(player, monsters.RandomMonster())
	view.PrintPlayerInfo(player)
}
