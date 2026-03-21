package game
import (
    "fmt"
    "math/rand"
)

func DiceRoll(sides int) int {
	 return rand.Intn(sides) + 1
}

func Roll20() int { return Roll(20) }
func Roll6() int  { return Roll(6) }