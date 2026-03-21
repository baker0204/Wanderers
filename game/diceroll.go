package game
import "math/rand"

// lol thanks Claude and https://programmingforlovers.com/chapter-2-forecasting-a-presidential-election-with-monte-carlo-simulation/chapter-2-go-code-along-randomized-algorithms-house-edge-of-craps-and-building-an-election-simulator/generating-random-integers-and-simulating-craps/
func DiceRoll(sides int) int {
	 return rand.Intn(sides) + 1 // Effective range is (0+1, sides+1)
}

func Roll20() int { return DiceRoll(20) } // (1,20)
func Roll6() int  { return DiceRoll(6) } // (1,6)