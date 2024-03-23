package utils

import (
	"fmt"
	"strings"
)

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

func (p Player) String() string {
	// use strings.Builder because += creates a copies of the the string
	var sb strings.Builder
	sb.WriteString("Player:\n")
	sb.WriteString(fmt.Sprintf("  ID: %d\n", p.ID))
	sb.WriteString(fmt.Sprintf("  Name: %s\n", p.name))
	sb.WriteString(fmt.Sprintf("  Chips: %d\n", p.chips))
	if len(p.Hand) >= 2 {
		sb.WriteString(fmt.Sprintf(" Hand: [%s, %s]\n", p.Hand[0].String(), p.Hand[1].String()))
	} else {
		sb.WriteString(" Hand: []\n")
	}

	return sb.String()
}
