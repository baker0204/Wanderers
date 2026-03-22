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

func (c *Character) Move(grid [][]Tile) bool {
    Direction := DiceRoll(4)
    newX := c.X
    newY := c.Y

    switch Direction {
    case 1:
        newY++
    case 2:
        newY--
    case 3:
        newX++
    case 4:
        newX--
    }

    if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
        if grid[newY][newX].Walkable {
            grid[c.Y][c.X].Entity = nil
            c.X = newX
            c.Y = newY
            grid[c.Y][c.X].Entity = c
            return true
        } else {
            return false
        }
    } else {
        return false
    }
}