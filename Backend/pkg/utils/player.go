package utils

type Status int

const (
	Check Status = iota
	Fold
	Raise
	Call
)

type Player struct {
	hand   []Card
	chips  int
	status Status
}
