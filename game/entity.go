package game

type Entity interface {
	GetSymbol() rune
    GetName() string
	GetPosition() (int, int)
	IsHostile() bool
}