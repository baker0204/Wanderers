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

    // Hide cursor for cleaner rendering
    fmt.Print("\033[?25l")
    // Restore cursor on exit
    defer fmt.Print("\033[?25h")

    // Clear screen once at the start
    fmt.Print("\033[2J")

    for {
        game.DrawGrid(grid, j, party)
        
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
