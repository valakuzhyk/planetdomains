package card

import "strings"

// Group of cards that can be accessed by index
type Group struct {
	Cards []Card
}

func (g Group) String() string {
	return strings.Join(StringList(g.Cards...), ", ")
}

func (g *Group) Insert(i int, c Card) {
	g.Cards = append(g.Cards[:i], append([]Card{c}, g.Cards[i:]...)...)
}

func (g *Group) Add(c ...Card) {
	g.Cards = append(g.Cards, c...)
}

// Take removes the ith card from the group
func (g *Group) Take(i int) Card {
	c := g.Cards[i]
	g.Cards = append(g.Cards[:i], g.Cards[i+1:]...)
	return c
}

func (g *Group) TakeAll() []Card {
	cards := g.Cards
	g.Cards = []Card{}
	return cards
}

// Peek at a card without taking it.
func (g *Group) Peek(i int) Card {
	return g.Cards[i]
}

func (g *Group) Len() int {
	return len(g.Cards)
}
