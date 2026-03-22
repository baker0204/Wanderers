package game

import "math/rand"

type Critter struct {
    Name      string
    Symbol    rune
    X         int
    Y         int
    Hostile   bool
    MoveCount int
	MoveStyle string
	Mood string
	Health int
}

func (crit *Critter) GetSymbol() rune { return crit.Symbol }
func (crit *Critter) GetName() string { return crit.Name }
func (crit *Critter) GetPosition() (int, int) { return crit.X, crit.Y }
func (crit *Critter) IsHostile() bool { return crit.Hostile }

func RandomCritter() *Critter {
		critters := []*Critter{
			{Name: "Tanuki",       Symbol: 'T',  Hostile: false, MoveStyle: "random",  Mood: "mischievous", Health: 5},
			{Name: "Slime",        Symbol: 's',  Hostile: true,  MoveStyle: "toward",  Mood: "hostile",     Health: 3},
			{Name: "Fox",          Symbol: 'F',  Hostile: false, MoveStyle: "toward",  Mood: "curious",     Health: 4},
			{Name: "Forest Bear",  Symbol: 'b',  Hostile: false, MoveStyle: "slow",    Mood: "calm",        Health: 10},
			{Name: "Wind Spirit",  Symbol: '≋',  Hostile: false, MoveStyle: "erratic", Mood: "playful",     Health: 2},
			{Name: "Lantern Ghost",Symbol: 'G',  Hostile: true,  MoveStyle: "slow",    Mood: "hostile",     Health: 4},
			{Name: "Angry Daruma", Symbol: 'O',  Hostile: true,  MoveStyle: "charge",  Mood: "hostile",     Health: 6},
		}
    return critters[rand.Intn(len(critters))]
}

func NewSootSprite() *Critter {
    return &Critter{
        Name:      "Soot Sprite",
        Symbol:    '*',
        Hostile:   false,
        MoveStyle: "random",
        Mood:      "cozy",
        Health:    1,
    }
}

func NewTreeSpirit() *Critter {
    return &Critter{
        Name:      "Tree Spirit",
        Symbol:    '↑',
        Hostile:   false,
        MoveStyle: "rooted",
        Mood:      "calm",
        Health:    8,
    }
}

func NewKappa() *Critter {
    return &Critter{
        Name:      "Kappa",
        Symbol:    'κ',
        Hostile:   true,
        MoveStyle: "toward",
        Mood:      "hostile",
        Health:    6,
    }
}

type CritterManager struct {}