package cardimpl

import (
	"testing"
)

func Test_deckImpl_DrawAll(t *testing.T) {
	d := deckImpl{cards: []Card{
		Mock{Name: "Card 1"},
		Mock{Name: "Card 2"},
		Mock{Name: "Card 3"},
	}}

	allCards := d.DrawAll()
	if len(allCards) != 3 {
		t.Fatal("Did not draw all cards")
	}

	if !d.IsEmpty() {
		t.Fatal("Deck is not empty after drawing all cards")
	}
}

func Test_deckImpl_Draw(t *testing.T) {
	d := deckImpl{cards: []Card{
		Mock{Name: "Card 1"},
		Mock{Name: "Card 2"},
		Mock{Name: "Card 3"},
	}}

	card := d.Draw()
	if card.GetName() != "Card 1" {
		t.Fatal("Didn't draw the right card")
	}

	if d.Len() != 2 {
		t.Fatal("Didn't remove a card after drawing")
	}
}

func Test_deckImpl_PlaceOnTop(t *testing.T) {
	d := deckImpl{cards: []Card{
		Mock{Name: "Card 1"},
		Mock{Name: "Card 2"},
		Mock{Name: "Card 3"},
	}}

	d.PlaceOnTop(Mock{Name: "Card 0"})
	if d.Len() != 4 {
		t.Fatal("Didn't add a card to the deck")
	}

	if d.cards[0].GetName() != "Card 0" {
		t.Fatal("Didn't place the card on the top")
	}
}

func Test_deckImpl_PlaceOnBottom(t *testing.T) {
	d := deckImpl{cards: []Card{
		Mock{Name: "Card 1"},
		Mock{Name: "Card 2"},
		Mock{Name: "Card 3"},
	}}

	d.PlaceOnBottom(Mock{Name: "Card 4"})
	if d.Len() != 4 {
		t.Fatal("Didn't add a card to the deck")
	}

	if d.cards[3].GetName() != "Card 4" {
		t.Fatal("Didn't place the card on the bottom")
	}
}
