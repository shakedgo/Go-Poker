package utils

type Table struct {
	deck    Deck
	players []Player
}

func InitTable() Table {
	deck := NewDeck()
	deck.Shuffle()
	return Table{deck: deck, players: []Player{}}
}

func (t *Table) AddPlayer(chips int) {
	t.players = append(t.players, Player{chips: chips})
}

func (t *Table) Deal(d *Deck) {
	if len(*d) < 2*len(t.players) {
		return
	}

	for range 2 {
		for _, player := range t.players {
			player.hand = append(player.hand, (*d)[0])
			*d = (*d)[1:]
		}
	}
}
