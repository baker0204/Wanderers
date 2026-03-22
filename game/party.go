package game

import "math/rand"

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
    MoveCount int
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

func (c *Character) Move(grid [][]Tile, centerX int, centerY int, grouped bool) bool {
    Direction := DiceRoll(4)
    newX := c.X
    newY := c.Y

    distX := centerX - c.X
    distY := centerY - c.Y
    farFromCenter := distX*distX + distY*distY > 9 // more than ~3 tiles away

    if c.Tired && DiceRoll(10) <= 5 { return false }
    if c.Spooked { grouped = false }

    if grouped && farFromCenter && DiceRoll(10) <= 7 {
        if centerX > c.X { newX++ } else if centerX < c.X { newX-- }
        if centerY > c.Y { newY++ } else if centerY < c.Y { newY-- }
    } else {
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
    }
    if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
        if grid[newY][newX].Walkable && grid[newY][newX].Entity == nil {
            grid[c.Y][c.X].Entity = nil
            c.X = newX
            c.Y = newY
            grid[c.Y][c.X].Entity = c
            c.MoveCount++
            if c.MoveCount >= 20 && DiceRoll(20) <= c.MoveCount-19 { c.Tired = true }
            return true
        } else {
            return false
        }
    } else {
        return false
    }
}

type Party struct {
    Members  []*Character
    Grouped  bool
}

func NewParty(grid [][]Tile) *Party {
    classes := []string{"Knight", "Bard", "Druid", "Wizard", "Ranger", "Herbalist"}
    party := []*Character{}

    rand.Shuffle(len(classes), func(i, j int) {
        classes[i], classes[j] = classes[j], classes[i]
    })

    for _, class := range classes[:3] {
        char := &Character{
            Name:  classNames[class],
            Class: class,
            Trait: classTraits[class],
            Health: 10,
            X: DiceRoll(len(grid[0])-1),
            Y: DiceRoll(len(grid)-1),
        }
        grid[char.Y][char.X].Entity = char
        party = append(party, char)
    }

    return &Party{
        Members: party,
        Grouped: true,
    }
}

func (p *Party) Center() (int, int) {
    totalX := 0
    totalY := 0
    // calclulate the average X, Y of the party
    for _, c := range p.Members {
        totalX += c.X
        totalY += c.Y
    }

    return totalX / len(p.Members), totalY / len(p.Members)
}