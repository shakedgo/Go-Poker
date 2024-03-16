package utils

import (
	"errors"
	"fmt"
	"strings"
)

var Tables []*Table

var tableIDCounter int = 1

type Table struct {
	ID       int
	Deck     Deck
	players  []*Player
	minEnter int
}

func InitTable() Table {
	deck := NewDeck()
	deck.Shuffle()
	table := Table{ID: tableIDCounter, Deck: deck, players: []*Player{}, minEnter: 500}
	tableIDCounter++
	Tables = append(Tables, &table)
	return table
}

func GetTableByID(id int) *Table {
	for _, table := range Tables {
		if table.ID == id {
			return table
		}
	}
	return nil
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
	// use strings.Builder because += creates a copies of the the string
	var sb strings.Builder
	sb.WriteString("Table:\n")
	sb.WriteString(fmt.Sprintf("  ID: %d\n", t.ID))
	sb.WriteString(fmt.Sprintf("  Minimum Entry: %d\n", t.minEnter))
	sb.WriteString("  Players:\n")
	for _, player := range t.players {
		sb.WriteString(fmt.Sprintf("    - %s (Chips: %d)\n", player.name, player.chips))
		if player.Hand != nil {
			sb.WriteString(fmt.Sprintf("      - Hand: [%s, %s]\n", player.Hand[0].String(), player.Hand[1].String()))
		}
	}
	// sb.WriteString(fmt.Sprintf("  Deck: \n    %v\n", t.Deck))
	sb.WriteString(fmt.Sprintf("  Num Cards in Deck: %d\n", len(t.Deck)))
	return sb.String()
}
