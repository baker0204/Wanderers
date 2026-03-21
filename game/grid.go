package game

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func DrawGrid() {
	GridSlice := [][]rune{
		{'.', '.', '.', '.', '.', '.', '.', '♣', '♣', '♣'},
		{'.', '.', '.', '.', '.', '.', '♣', '♣', '♣', '♣'},
		{'.', '.', '.', '.', '.', '.', '♣', '♣', '♣', '♣'},
		{'.', '.', '.', '.', '.', '.', '.', '♣', '♣', '♣'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '♣', '♣'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '♣', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}

	// Each cell is "X " = 2 chars wide, so grid width = cols * 2
	gridCols := len(GridSlice[0])
	gridRows := len(GridSlice)
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

	// Print the grid
	for y := 0; y < gridRows; y++ {
		fmt.Print(padding)
		for x := 0; x < gridCols; x++ {
			fmt.Printf("%c ", GridSlice[y][x]) // space after each cell
		}
		fmt.Println()
	}
}