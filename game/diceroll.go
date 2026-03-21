package game
import "math/rand"

func DiceRoll(sides int) int {
	 return rand.Intn(sides) + 1
}

func Roll20() int { return DiceRoll(20) }
func Roll6() int  { return DiceRoll(6) }