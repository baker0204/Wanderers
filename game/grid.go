package game

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func GenerateGrid(x int, y int) [][]Tile {
	grid := [][]Tile{} 

	for a := 0; a < y; a++ {
		row := []Tile{}
		for b := 0; b < x; b++ {
			row = append(row, Tile{Symbol: '.', Walkable: true, Terrain: "plains"})
		}
		grid = append(grid, row)
	}

	return grid
}

func DrawGrid(grid [][]Tile) {
	// Each cell is "X " = 2 chars wide, so grid width = cols * 2
	gridCols := len(grid[0])
	gridRows := len(grid)
	cellWidth := 2
	gridWidth := gridCols * cellWidth

	// Get terminal width, fall back to 80 if unavailable
	termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		termWidth = 80
	}

	// Calculate left padding to center the grid
	leftPad := (termWidth - gridWidth) / 2
	if leftPad < 0 {
		leftPad = 0
	}
	padding := fmt.Sprintf("%*s", leftPad, "")

	// Vertical padding — push grid down a few lines
	termHeight, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		termHeight = 24
	}
	topPad := (termHeight - gridRows) / 2
	for i := 0; i < topPad; i++ {
		fmt.Println()
	}

	// "Draw" the grid
	for y := 0; y < gridRows; y++ {
		fmt.Print(padding)
		for x := 0; x < gridCols; x++ {
			//fmt.Printf("%c ", grid[y][x]) // space after each cell
			if grid[y][x].Entity != nil { 
				fmt.Printf("%c ", grid[y][x].Entity.GetSymbol()) 
			} else { fmt.Printf("%c ", grid[y][x].Symbol) }
		}
		fmt.Println()
	}
}