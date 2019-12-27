package characters

import (
	"fmt"
	"fun_stuff/models/items"
	"fun_stuff/models/monsters"
)

type (
	genericCharacter struct {
		_type     string
		_adj      string
		_deity    string
		Name      string
		Hp        int64
		BaseDmg   int64
		BaseArmor int64
		Equipment []items.Item
	}

	Character interface {
		Attack(*monsters.Monster)
		Armor() int64
		Dmg() int64
		Fight(*monsters.Monster)
		Equip(items.Item)
	}
)

func (c *genericCharacter) Equip(item items.Item) {
	fmt.Printf("%s equips a %s %s!\n", c.Name, item.Rarity, item.Name)
	c.Equipment = append(c.Equipment, item)
}

func (c *genericCharacter) Fight(m *monsters.Monster) {
	fmt.Printf("%s is fighting a filthy %s! The battle unravels...\n", c.Name, m.Name)
	for c.Hp > 0 && m.Hp > 0 {
		c.Attack(m)
	}
	if c.Hp <= 0 {
		fmt.Printf("By the gods! A %s %s has been slain by the filthy %s!\n", c._adj, c.Name, m.Name)
	} else {
		fmt.Printf("Praise the %s! A filthy %s perishes!\n", c._deity, m.Name)
	}
}

func (c *genericCharacter) Armor() int64 {
	var final int64 = 0
	for _, item := range c.Equipment {
		final += item.Armor
	}
	return final + c.BaseArmor
}

func (c *genericCharacter) Dmg() int64 {
	var final int64 = 0
	for _, item := range c.Equipment {
		final += item.Dmg
	}
	return final + c.BaseDmg
}

func (c *genericCharacter) Attack(m *monsters.Monster) {
	cDmg := c.Dmg() - m.Armor
	m.Hp -= cDmg
	fmt.Printf("%s attacks for %d! ", c.Name, cDmg)
	if m.Hp <= 0 {
		return
	}
	mDmg := m.Dmg - c.Armor()
	c.Hp -= mDmg
	fmt.Printf("A %s fights back, dealing %d damage!\n", m.Name, mDmg)
}

func (c *genericCharacter) String() string {
	str := fmt.Sprintf("A %s %s - %s!", c._adj, c._type, c.Name)
	if c.Hp > 0 {
		str += fmt.Sprintf(" As of now, he has %d HP, %d armor and deals %d damage per hit.\n", c.Hp, c.Armor(), c.Dmg())
		if len(c.Equipment) > 0 {
			str += "Equipment:\n"
			for _, item := range c.Equipment {
				str += " > " + fmt.Sprintln(item)
			}
		}
	} else {
		str += " Apparently dead."
	}
	return str
}