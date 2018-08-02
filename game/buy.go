package game

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/valakuzhyk/planetdomains/card"
)

func purchaseCards(p *person) {
	// If possible try and purchase some cards
	// 0 represents explorer
	var i int
	for true {
		log.Debug(p.field)
		printPurchaseOptions(p.field)
		log.Printf("You have %d Trade, would you like to buy a card?", p.GetTrade())
		fmt.Scanf("%d", &i)
		if i < 0 {
			break
		}

		if i == 0 {
			buyExplorer(p)
		} else {
			buyFromTradeRow(p, i)
		}
	}
	return
}

func printPurchaseOptions(field *Field) {
	log.Printf("Cards for Purchase:")
	if field.Explorers.IsEmpty() {
		log.Printf("   0: N/A XXXX")
	} else {
		log.Printf("   0: %s %d", field.Explorers.Peek().GetName(), field.Explorers.Peek().GetCost())
	}
	for i, c := range field.TradeRow {
		log.Printf("   %d: %s %d", i+1, card.String(c), c.GetCost())
	}
}

func buyExplorer(p *person) {
	explorers := p.field.Explorers
	if explorers.IsEmpty() {
		log.Println("No more explorers!")
		return
	}
	eCost := explorers.Peek().GetCost()
	if eCost > p.GetTrade() {
		log.Printf("Cost %d greater than your trade %d. Insufficient Funds",
			eCost,
			p.GetTrade())
		return
	}
	e := explorers.Draw()
	p.AddTrade(-e.GetCost())
	p.Discard.PlaceOnTop(e)
}

func buyFromTradeRow(p *person, n int) {
	tradeRow := p.field.TradeRow
	if n > len(tradeRow) {
		log.Println("Invalid choice")
		return
	}

	// User input is 1-indexed
	c := tradeRow[n-1]

	if p.GetTrade() < c.GetCost() {
		log.Printf("Cost %d greater than your trade %d. Insufficient Funds",
			c.GetCost(),
			p.GetTrade())
		return
	}

	p.AddTrade(-c.GetCost())
	p.field.TradeRow = append(tradeRow[:n-1], tradeRow[n:]...)

	// Replace the card on the trade row if possible
	if !p.field.TradeDeck.IsEmpty() {
		p.field.TradeRow = append(p.field.TradeRow, p.field.TradeDeck.Draw())
	}

	p.Discard.PlaceOnTop(c)
}
