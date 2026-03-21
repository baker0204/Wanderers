package game

//import "fmt"

type Character struct {
    Name      string
    Class     string
    Trait     string
    Health    int
    Tired     bool
    Spooked   bool
}

var classTraits = map[string]string{
    "Knight":    "Overcautious",
    "Bard":      "Dramatic",
    "Druid":     "Easily Distracted",
    "Wizard":    "Forgetful",
    "Ranger":    "Suspicious",
    "Herbalist": "Optimistic",
}

var classNames = map[string]string{
    "Knight":    "Elra",
    "Bard":      "Pip",
    "Druid":     "Moss",
    "Wizard":    "Aldric",
    "Ranger":    "Syla",
    "Herbalist": "Basil",
}