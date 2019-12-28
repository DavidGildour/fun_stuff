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
	player.Equip(items.Sword(rarity))
	view.PrintPlayerInfo(player)
}

func FightRandomMonster() {
	player.Fight(monsters.RandomMonster())
	view.PrintPlayerInfo(player)
}
