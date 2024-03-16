package utils

import (
	"errors"
	"fmt"
	"strings"
)

type Table struct {
	Deck     Deck
	players  []*Player
	minEnter int
}

func InitTable() Table {
	deck := NewDeck()
	deck.Shuffle()
	return Table{Deck: deck, players: []*Player{}, minEnter: 500}
}

func (t *Table) JoinTable(player *Player) error {
	if player.chips < t.minEnter {
		return errors.New("not enough chips")
	} else {
		t.players = append(t.players, player)
		return nil
	}
}

func (t *Table) Deal(d *Deck) {
	if len(*d) < 2*len(t.players) {
		return
	}

	for range 2 {
		for _, player := range t.players {
			player.Hand = append(player.Hand, (*d)[0])
			*d = (*d)[1:]
		}
	}
}

func (t Table) String() string {
	// var tableString string
	// tableString += "Table:\n"
	// tableString += fmt.Sprintf("  Minimum Entry: %d\n", t.minEnter)
	// tableString += "  Players:\n"
	// for _, player := range t.players {
	// 	tableString += fmt.Sprintf("    - %s (Chips: %d)\n", player.name, player.chips)
	// }
	// tableString += fmt.Sprintf("  Deck: \n    %v\n", t.Deck)
	// return tableString

	// use strings.Builder because += creates a copies of the the string
	var sb strings.Builder
	sb.WriteString("Table:\n")
	sb.WriteString(fmt.Sprintf("  Minimum Entry: %d\n", t.minEnter))
	sb.WriteString("  Players:\n")
	for _, player := range t.players {
		sb.WriteString(fmt.Sprintf("    - %s (Chips: %d)\n", player.name, player.chips))
		sb.WriteString(fmt.Sprintf("    -  Hand: [%s, %s]\n", player.Hand[0].String(), player.Hand[1].String()))
	}
	// sb.WriteString(fmt.Sprintf("  Deck: \n    %v\n", t.Deck))
	sb.WriteString(fmt.Sprintf("  Num Cards in Deck: %d\n", len(t.Deck)))
	return sb.String()
}
