package internal

// Player represents what operations can be done on a player by
// the game mechanics. This cannot have any operations on card.
type Player interface {
	AddTrade(int) error
	AddCombat(int) error
	AddAuthority(int)
	MustDiscard()
	DrawCards(uint)
	DestroyBase(p Player)
	ScrapFromHand()
	ScrapFromDiscard()
	ScrapFromHandOrDiscard(uint)
	ScrapFromTradeRow(uint)
}
