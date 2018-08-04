package game

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func purchaseCards(p *person) {
	// If possible try and purchase some cards
	// 0 represents explorer
	for true {
		log.Debug(p.field)
		printPurchaseOptions(p.field)
		fmt.Printf("You have %d Trade, would you like to buy a card?\n", p.GetTrade())
		var i int
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			fmt.Printf("Could not understand input. Try again\n")
			continue
		}
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
	fmt.Printf("Cards for Purchase:\n")
	if field.Explorers.IsEmpty() {
		fmt.Printf("   0: N/A XXXX\n")
	} else {
		fmt.Printf("   0: %s %d\n", field.Explorers.Peek().GetName(), field.Explorers.Peek().GetCost())
	}
	for i, c := range field.TradeRow {
		fmt.Printf("   %d: %s %d\n", i+1, c.String(), c.GetCost())
	}
}

func buyExplorer(p *person) {
	explorers := p.field.Explorers
	if explorers.IsEmpty() {
		fmt.Println("No more explorers!\n")
		return
	}
	eCost := explorers.Peek().GetCost()
	if eCost > p.GetTrade() {
		fmt.Printf("Cost %d greater than your trade %d. Insufficient Funds\n",
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
		fmt.Println("Invalid choice\n")
		return
	}

	// User input is 1-indexed
	c := tradeRow[n-1]

	if p.GetTrade() < c.GetCost() {
		fmt.Printf("Cost %d greater than your trade %d. Insufficient Funds\n",
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
