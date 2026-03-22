package game

import "math/rand"

type Event struct {
    Description string
    Difficulty  int
}

var dangerEvents = []Event{
    {"A grumpy River Troll blocks the path", 12},
    {"A thunderstorm rolls in suddenly", 8},
    {"A wolf pup won't stop following the party", 6},
    {"A mischievous wind spirit steals someone's hat", 10},
    {"A very large, very lost bear sits in the road", 14},
}

var calmEvents = []Event{
    {"A small patch of glowing flowers appears", 0},
    {"Someone finds a warm loaf of bread on a stump", 0},
    {"A friendly crow offers directions (maybe)", 0},
}

func RandomEvent(difficulty int) Event {
    if DiceRoll(10) <= difficulty {
        return dangerEvents[rand.Intn(len(dangerEvents))]
    }
    return calmEvents[rand.Intn(len(calmEvents))]
}