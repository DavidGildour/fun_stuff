package characters

type Profession func(string) Character

var Professions = map[string]Profession{
	"1": Warrior,
	"2": Huntress,
}
