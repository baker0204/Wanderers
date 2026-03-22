package game

import (
	"fmt"
	"os"
	"strings"
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

const clearLine = "\033[2K\r"

// I'm gonna be honest Claude did most of this UI code
func DrawGrid(grid [][]Tile, j *Journey) {
	gridCols := len(grid[0])
	gridRows := len(grid)
	cellWidth := 2
	gridWidth := gridCols * cellWidth

	// Get terminal dimensions
	termWidth, termHeight, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		termWidth = 80
		termHeight = 24
	}

	// Center horizontally
	leftPad := (termWidth - gridWidth) / 2
	if leftPad < 0 {
		leftPad = 0
	}
	padding := strings.Repeat(" ", leftPad)

	// Status bar
	phase := j.Phases[j.Phase]
	day := fmt.Sprintf("Day %d", j.Daycount)
	statusBar := fmt.Sprintf("%s  |  %s", phase, day)

	// Legend
	legend := ". plains  ♣ forest  ~ water  K knight  B bard  D druid  W wizard  R ranger  H herbalist"
	legendPad := (termWidth - len([]rune(legend))) / 2
	if legendPad < 0 {
		legendPad = 0
	}
	fmt.Print(clearLine)
	fmt.Print(strings.Repeat(" ", legendPad) + legend + "\n")

	// Center vertically as a block
	totalLines := gridRows + 3
	topPad := (termHeight - totalLines) / 2
	if topPad < 0 {
		topPad = 0
	}

	// Move cursor to top-left
	fmt.Print("\033[H")

	// Top padding — clear each line
	for i := 0; i < topPad; i++ {
		fmt.Print(clearLine + "\n")
	}

	// Title
	title := "Wanderers 🌿"
	titlePad := (termWidth - len([]rune(title))) / 2
	if titlePad < 0 {
		titlePad = 0
	}
	fmt.Print(clearLine)
	fmt.Print(strings.Repeat(" ", titlePad) + title + "\n")

	// Draw the grid — clear each line before writing
	for y := 0; y < gridRows; y++ {
		fmt.Print(clearLine)
		fmt.Print(padding)
		for x := 0; x < gridCols; x++ {
			if grid[y][x].Entity != nil {
				fmt.Printf("%c ", grid[y][x].Entity.GetSymbol())
			} else {
				fmt.Printf("%c ", grid[y][x].Symbol)
			}
		}
		fmt.Print("\n")
	}

	// Blank line + status bar
	fmt.Print(clearLine + "\n")
	fmt.Print(clearLine)
	statusPad := (termWidth - len([]rune(statusBar))) / 2
	if statusPad < 0 {
		statusPad = 0
	}
	fmt.Print(strings.Repeat(" ", statusPad) + statusBar + "\n")
}