package game

import (
    "fmt"
    "math"
    "math/rand"
)

type Critter struct {
    Name      string
    Symbol    rune
    X         int
    Y         int
    Hostile   bool
    MoveCount int
	MoveStyle string
	Mood      string
	Health    int
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

type CritterManager struct {
    Critters []*Critter
}

func NewCritterManager() *CritterManager {
    return &CritterManager{
        Critters: []*Critter{},
    }
}

func (cm *CritterManager) Spawn(critter *Critter, grid [][]Tile) {
    for {
        x := DiceRoll(len(grid[0])-1)
        y := DiceRoll(len(grid)-1)
        if grid[y][x].Walkable && grid[y][x].Entity == nil {
            critter.X = x  // ← in here
            critter.Y = y  // ← in here
            grid[y][x].Entity = critter
            cm.Critters = append(cm.Critters, critter)
            break
        }
    }
}

func (cm *CritterManager) MoveAll(grid [][]Tile, party *Party) {
    for _, crit := range cm.Critters {
        // move crit based on its MoveStyle
        Direction := DiceRoll(4)
        newX := crit.X
        newY := crit.Y
        switch crit.MoveStyle {
            case "random":
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
            case "toward":
                    cx, cy := party.Center()
                    if cx > crit.X { newX++ } else if cx < crit.X { newX-- }
                    if cy > crit.Y { newY++ } else if cy < crit.Y { newY-- }
            case "slow":
                if DiceRoll(2) == 1 {break} else {
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
            case "erratic":
                Direction2 := DiceRoll(4)
                if DiceRoll(2) == 1 {
                    Direction = Direction2
                }
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

            case "rooted":
                continue
            case "charge":
                closest := party.Members[0]
                closestDist := math.MaxInt32
                for _, member := range party.Members {
                    distX := member.X - crit.X
                    distY := member.Y - crit.Y
                    dist := distX*distX + distY*distY
                    if dist < closestDist {
                        closestDist = dist
                        closest = member
                    }
                }
                if closest.X > crit.X { newX++ } else if closest.X < crit.X { newX-- }
                if closest.Y > crit.Y { newY++ } else if closest.Y < crit.Y { newY-- }
        }
            if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
        if grid[newY][newX].Walkable && grid[newY][newX].Entity == nil {
            grid[crit.Y][crit.X].Entity = nil
            crit.X = newX
            crit.Y = newY
            grid[crit.Y][crit.X].Entity = crit
        }
    }
    }  // ← for loop closes here
}


func (cm *CritterManager) CheckAdjacency(party *Party) []string {
    messages := []string{}
    for _, member := range party.Members {
        for _, crit := range cm.Critters {
            distX := crit.X - member.X
            distY := crit.Y - member.Y
            if distX < 0 { distX = -distX } // absolute value
            if distY < 0 { distY = -distY }
            adjacent := distX <= 1 && distY <= 1

            if adjacent {
            if crit.Hostile {
                messages = append(messages, fmt.Sprintf("⚔️ %s attacks %s!", crit.Name, member.Name))
            } else {
                messages = append(messages, fmt.Sprintf("✨ %s encounters %s!", member.Name, crit.Name))
            }
            }
        }
    }
    return messages
}