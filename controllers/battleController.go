package controllers

import (
	"fun_stuff/models/characters"
	"fun_stuff/models/monsters"
	"fun_stuff/view"
)

var fightChoices = []string{
	"Attack",
}

// Fight ...
func Fight(c characters.Character, m *monsters.Monster) {
	view.AnnounceBattle(c.Name(), m.Name)
	for c.Hp() > 0 && m.Hp > 0 {
		view.GetPlayerActionChoice(fightChoices)
		c.Attack(m)
	}
	if c.Hp() <= 0 {
		view.AnnounceDefeat(c, m.Name)
	} else {
		view.AnnounceVictory(c, m.Name)
	}
}
