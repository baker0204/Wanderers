package game

//import "fmt"

type Character struct {
    Name      string
    Class     string
    Trait     string
    Health    int
    Tired     bool
    Spooked   bool
    Symbol    rune
    X         int
    Y         int
    Hostile   bool
}

var classTraits = map[string]string{
    "Knight":    "Overcautious",
    "Bard":      "Dramatic",
    "Druid":     "Easily Distracted",
    "Wizard":    "Forgetful",
    "Ranger":    "Suspicious",
    "Herbalist": "Optimistic",
}

var classNames = map[string]string{
    "Knight":    "Elra",
    "Bard":      "Pip",
    "Druid":     "Moss",
    "Wizard":    "Aldric",
    "Ranger":    "Rowan",
    "Herbalist": "Basil",
}

func (c *Character) GetSymbol() rune {
    switch c.Class {
    case "Knight":
        return 'K'
    case "Bard":
        return 'B'
    case "Druid":
        return 'D'
    case "Wizard":
        return 'W'
    case "Ranger":
        return 'R'
    case "Herbalist":
        return 'H'
    default:
        return '?'
    }
}

func (c *Character) GetName() string { return c.Name }

func (c *Character) GetPosition() (int, int) { return c.X, c.Y }

func (c *Character) IsHostile() bool { return false }