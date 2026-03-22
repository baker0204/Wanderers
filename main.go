package main

import (
    "fmt"
    "time"
    "wanderers/game"
)

func main() {
    grid := game.GenerateGrid(30, 20)

    knight := &game.Character{Name: "Elra", Class: "Knight", X: 5, Y: 5}
    grid[knight.Y][knight.X].Entity = knight

    j := game.NewJourney()

    // Hide cursor for cleaner rendering
    fmt.Print("\033[?25l")
    // Restore cursor on exit
    defer fmt.Print("\033[?25h")

    // Clear screen once at the start
    fmt.Print("\033[2J")

    for {
        game.DrawGrid(grid, j)
        knight.Move(grid)
        j.Tick()
        time.Sleep(200 * time.Millisecond)
    }
}