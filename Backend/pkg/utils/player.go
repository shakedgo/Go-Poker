package utils

type Status int

const (
	Check Status = iota
	Fold
	Raise
	Call
)

type Player struct {
	name   string
	Hand   []Card
	chips  int
	status Status
}

func NewPlayer(name string) Player {
	return Player{name: name, chips: 1000}
}
