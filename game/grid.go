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

	// Pass 1 — seed randomly
	for y := range grid {
		for x := range grid[y] {
			roll := DiceRoll(18)
			if roll == 1 {
				grid[y][x] = Tile{Symbol: '~', Walkable: false, Terrain: "water"}
			} else if roll == 2 {
				grid[y][x] = Tile{Symbol: '♣', Walkable: false, Terrain: "forest"}
			}
		}
	}

	// Pass 2 — smooth by spreading terrain to neighbors
	for y := range grid {
		for x := range grid[y] {
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dx == 0 && dy == 0 { continue }
					if y+dy < 0 || y+dy >= len(grid) || x+dx < 0 || x+dx >= len(grid[0]) { continue } // bounds check
					if grid[y+dy][x+dx].Terrain == "water" {
						roll := DiceRoll(10)
						if roll <= 2 {
							grid[y][x] = Tile{Symbol: '~', Walkable: false, Terrain: "water"}
						}
					}
					if grid[y+dy][x+dx].Terrain == "forest" {
						roll := DiceRoll(10)
						if roll <= 2 {
							grid[y][x] = Tile{Symbol: '♣', Walkable: false, Terrain: "forest"}
						}
					}
				}
			}
		}
	}

	return grid
}

const clearLine = "\033[2K\r"

// visWidth returns the real terminal display width of a string,
// accounting for double-wide emoji characters
func visWidth(s string) int {
	width := 0
	for _, r := range s {
		if r >= 0x2300 {
			width += 2 // covers ✅ and all emoji
		} else {
			width++
		}
	}
	return width
}

// padTo pads a string to exactly targetWidth visible columns
func padTo(s string, targetWidth int) string {
	vis := visWidth(s)
	if vis >= targetWidth {
		return s
	}
	return s + strings.Repeat(" ", targetWidth-vis)
}

// buildPanel builds the status panel lines including the box frame
func buildPanel(p *Party, panelWidth int) []string {
	inner := panelWidth - 2

	lines := []string{}
	lines = append(lines, "┌"+strings.Repeat("─", inner)+"┐")

	for i, c := range p.Members {
		nameStr := fmt.Sprintf("%c  %s the %s", c.GetSymbol(), c.Name, c.Class)
		lines = append(lines, "│"+padTo(" "+nameStr, inner)+"│")

		status := ""
		if c.Tired   { status += "😴 Tired  " }
		if c.Spooked { status += "😨 Spooked" }
		if status == "" { status = "✅ OK" }
		lines = append(lines, "│"+padTo("  "+status, inner)+"│")

		if i < len(p.Members)-1 {
			lines = append(lines, "│"+strings.Repeat(" ", inner)+"│")
		}
	}

	lines = append(lines, "└"+strings.Repeat("─", inner)+"┘")
	return lines
}

// buildEventLog builds the event log panel with tilde borders
func buildEventLog(messages []string, logWidth int) []string {
	inner := logWidth - 2 // space between the two ~ borders

	lines := []string{}

	// Header
	lines = append(lines, strings.Repeat("~", logWidth))
	lines = append(lines, "~"+padTo(" Event Log", inner)+"~")
	lines = append(lines, strings.Repeat("~", logWidth))

	// Pad to 5 entries so height is always consistent
	entries := make([]string, 5)
	copy(entries, messages)

	for _, msg := range entries {
		if msg == "" {
			lines = append(lines, "~"+strings.Repeat(" ", inner)+"~")
		} else {
			// truncate to inner width respecting emoji double-width
			display := " " + msg
			width := 0
			truncated := ""
			for _, r := range display {
				charW := 1
				if r >= 0x2300 {
					charW = 2
				}
				if width+charW > inner {
					break
				}
				truncated += string(r)
				width += charW
			}
			lines = append(lines, "~"+padTo(truncated, inner)+"~")
		}
	}

	// Footer
	lines = append(lines, strings.Repeat("~", logWidth))
	return lines
}

// I'm gonna be honest Claude did most of this UI code
func DrawGrid(grid [][]Tile, j *Journey, p *Party, messages []string) {
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

	panelWidth := 28
	logWidth   := 28
	gap := 2

	// Total width = log + gap + grid + gap + panel
	totalWidth := logWidth + gap + gridWidth + gap + panelWidth

	// Center horizontally
	leftPad := (termWidth - totalWidth) / 2
	if leftPad < 0 {
		leftPad = 0
	}
	padding := strings.Repeat(" ", leftPad)

	// Build panel and log lines
	panel    := buildPanel(p, panelWidth)
	eventLog := buildEventLog(messages, logWidth)

	// Status bar
	phase := j.Phases[j.Phase]
	day := fmt.Sprintf("Day %d", j.Daycount)
	statusBar := fmt.Sprintf("%s  |  %s", phase, day)

	// Legend
	legend := ". Plains  ♣ Forest  ~ Water  K Knight  B Bard  D Druid  W Wizard  R Ranger  H Herbalist"
	legendPad := (termWidth - len([]rune(legend))) / 2
	if legendPad < 0 {
		legendPad = 0
	}

	// Total lines: title + grid + blank + statusbar + blank + legend
	totalLines := 1 + gridRows + 1 + 1 + 1 + 1
	topPad := (termHeight - totalLines) / 2
	if topPad < 0 {
		topPad = 0
	}

	// Move cursor to top-left
	fmt.Print("\033[H")

	// Top padding
	for i := 0; i < topPad; i++ {
		fmt.Print(clearLine + "\n")
	}

	// Title
	title := "Wanderers 🌿"
	titlePad := (termWidth - visWidth(title)) / 2
	if titlePad < 0 {
		titlePad = 0
	}
	fmt.Print(clearLine)
	fmt.Print(strings.Repeat(" ", titlePad) + title + "\n")

	// Draw grid rows with log on left and panel on right
	gapStr := strings.Repeat(" ", gap)
	for y := 0; y < gridRows; y++ {
		fmt.Print(clearLine)
		fmt.Print(padding)

		// Event log column
		if y < len(eventLog) {
			fmt.Print(eventLog[y])
		} else {
			fmt.Print(strings.Repeat(" ", logWidth))
		}

		fmt.Print(gapStr)

		// Grid
		for x := 0; x < gridCols; x++ {
			if grid[y][x].Entity != nil {
				fmt.Printf("%c ", grid[y][x].Entity.GetSymbol())
			} else {
				fmt.Printf("%c ", grid[y][x].Symbol)
			}
		}

		// Status panel
		fmt.Print(gapStr)
		if y < len(panel) {
			fmt.Print(panel[y])
		}
		fmt.Print("\n")
	}

	// Blank line + centered status bar
	fmt.Print(clearLine + "\n")
	fmt.Print(clearLine)
	statusPad := (termWidth - visWidth(statusBar)) / 2
	if statusPad < 0 {
		statusPad = 0
	}
	fmt.Print(strings.Repeat(" ", statusPad) + statusBar + "\n")

	// Blank line + legend
	fmt.Print(clearLine + "\n")
	fmt.Print(clearLine)
	fmt.Print(strings.Repeat(" ", legendPad) + legend + "\n")
}
