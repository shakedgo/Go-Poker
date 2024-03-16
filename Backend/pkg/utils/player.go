package utils

type Status int

const (
	Check Status = iota
	Fold
	Raise
	Call
)

var Players []*Player

var playerIDCounter int = 1

type Player struct {
	ID     int
	name   string
	Hand   []Card
	chips  int
	status Status
}

func NewPlayer(name string) Player {
	player := Player{ID: playerIDCounter, name: name, chips: 1000}
	playerIDCounter++
	Players = append(Players, &player)
	return player
}

func GetPlayerByID(id int) *Player {
	for _, player := range Players {
		if player.ID == id {
			return player
		}
	}
	return nil
}
