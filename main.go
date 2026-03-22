package main

import (
    "fmt"
    "wanderers/game"
    "time"
)

func main(){
    grid := game.GenerateGrid(30, 20)
    
    // Knight Test Case
    knight := &game.Character{Name: "Elra", Class: "Knight", X: 5, Y: 5}
    grid[5][5].Entity = knight

    game.DrawGrid(grid)

    grid[knight.Y][knight.X].Entity = knight
    j := game.NewJourney()

    // test game loop
    for {
        fmt.Print("\033[H\033[2J")
        game.DrawGrid(grid)
        knight.Move(grid)
        j.Tick()
        fmt.Println(j.Phases[j.Phase])
        time.Sleep(200 * time.Millisecond)
    }

    fmt.Print("Wanderers 🌿: ")
    fmt.Println("A Fantasy Adventure Simulation Game")
}