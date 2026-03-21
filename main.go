package main

import (
    "fmt"
    "wanderers/game"
)

func main(){
    grid := game.GenerateGrid(30, 20)
    knight := &game.Character{Name: "Elra", Class: "Knight", X: 5, Y: 5}
    grid[5][5].Entity = knight
    game.DrawGrid(grid)
    fmt.Print("Wanderers 🌿: ")
    fmt.Println("A Fantasy Adventure Simulation Game")
}