package main

import (
    "fmt"
    "time"
    "wanderers/game"
)

func main() {
    grid := game.GenerateGrid(30, 20)
    party := game.NewParty(grid)
    j := game.NewJourney()
    cm := game.NewCritterManager()

    // Hide cursor for cleaner rendering
    fmt.Print("\033[?25l")
    // Restore cursor on exit
    defer fmt.Print("\033[?25h")

    // Clear screen once at the start
    fmt.Print("\033[2J")

    cm.Spawn(game.RandomCritter(), grid)
    cm.Spawn(game.RandomCritter(), grid)
    cm.Spawn(game.RandomCritter(), grid)

    eventLog := []string{}

    for {
        game.DrawGrid(grid, j, party, eventLog)
        cm.MoveAll(grid, party)
        messages := cm.CheckAdjacency(party)
        for _, msg := range messages {
            eventLog = append(eventLog, msg)
            if len(eventLog) > 5 {
                eventLog = eventLog[len(eventLog)-5:]
            }
        }
        
        cx, cy := party.Center()
        for _, character := range party.Members {
            character.Move(grid, cx, cy, party.Grouped)
        }
        if j.Phases[j.Phase] == "💤 Rest" {
            for _, c := range party.Members {
                c.Tired = false
                c.Spooked = false
                c.MoveCount = 0
            }
        }
        
        if j.Phases[j.Phase] == "🌅 Morning" && j.CurrentTick == 0 {
            if game.DiceRoll(7) == 7 {
                cx, cy := party.Center()
                sprite := game.NewSootSprite()
                // find an empty tile near the center
                for _, offset := range [][2]int{{0,1},{0,-1},{1,0},{-1,0}} {
                    nx, ny := cx+offset[0], cy+offset[1]
                    if grid[ny][nx].Walkable && grid[ny][nx].Entity == nil {
                        sprite.X = nx
                        sprite.Y = ny
                        grid[ny][nx].Entity = sprite
                        cm.Critters = append(cm.Critters, sprite)
                        break
                    }
                }
            }
        }

        j.Tick()
        time.Sleep(200 * time.Millisecond)
    }
}
