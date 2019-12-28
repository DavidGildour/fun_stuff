package characters

import (
	"fmt"
	"fun_stuff/models/items"
	"fun_stuff/models/monsters"
)

type (
	// genericCharacter is a generc character
	genericCharacter struct {
		_type     string
		_adj      string
		_deity    string
		name      string
		hp        int64
		baseDmg   int64
		baseArmor int64
		equipment []items.Item
	}

	// Character is an interface for interact with a character structs
	Character interface {
		Adj() string
		Type() string
		Deity() string
		Name() string
		Hp() int64
		Equipment() []items.Item
		Attack(*monsters.Monster)
		Armor() int64
		Dmg() int64
		Equip(items.Item)
	}
)

// Type returns character type
func (c *genericCharacter) Type() string {
	return c._type
}

// Deity returns character deity
func (c *genericCharacter) Deity() string {
	return c._deity
}

// Adj returns character adjective
func (c *genericCharacter) Adj() string {
	return c._adj
}

// Name returns character name
func (c *genericCharacter) Name() string {
	return c.name
}

// Hp returns character hp
func (c *genericCharacter) Hp() int64 {
	return c.hp
}

// Equipment returns quipment
func (c *genericCharacter) Equipment() []items.Item {
	return c.equipment
}

// Equip is a method for equipping new item
func (c *genericCharacter) Equip(item items.Item) {
	c.equipment = append(c.equipment, item)
}

// Armor is a method for summing up and returning the charater armor based on ever equipmnt it has.
func (c *genericCharacter) Armor() int64 {
	var final int64 = 0
	for _, item := range c.equipment {
		final += item.Armor
	}
	return final + c.baseArmor
}

// Dmg is a method for returning sum of all dmg character has based on equipment.
func (c *genericCharacter) Dmg() int64 {
	var final int64 = 0
	for _, item := range c.equipment {
		final += item.Dmg
	}
	return final + c.baseDmg
}

// Attack is a method for dealing damage to a monster.
func (c *genericCharacter) Attack(m *monsters.Monster) {
	cDmg := c.Dmg() - m.Armor
	m.Hp -= cDmg
	fmt.Printf("%s attacks for %d! ", c.Name(), cDmg)
	if m.Hp <= 0 {
		return
	}
	mDmg := m.Dmg - c.Armor()
	c.hp -= mDmg
	fmt.Printf("A %s fights back, dealing %d damage!\n", m.Name, mDmg)
}
