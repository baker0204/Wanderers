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

    for {
        game.DrawGrid(grid, j, party)
        cm.MoveAll(grid, party)
        messages := cm.CheckAdjacency(party)
        for _, msg := range messages {
            fmt.Println(msg)
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

        j.Tick()
        time.Sleep(200 * time.Millisecond)
    }
}
