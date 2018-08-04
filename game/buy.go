package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/valakuzhyk/planetdomains/card"

	"github.com/buger/goterm"
)

func purchaseCards(p *person) {
	// If possible try and purchase some cards
	// 0 represents explorer
	for true {
		goterm.Clear()
		goterm.Flush()
		choice := getPurchaseChoice(p.field)

		if choice < 0 {
			break
		}
		if choice == 0 {
			buyExplorer(p)
		} else {
			buyFromTradeRow(p, choice)
		}
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
	return
}

func getPurchaseChoice(field *Field) int {

	cardsToSelectFrom := field.TradeRow
	if !field.Explorers.IsEmpty() {
		cardsToSelectFrom = append([]card.Card{field.Explorers.Peek()}, cardsToSelectFrom...)
	}

	i := pickCard(cardsToSelectFrom)
	return i
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
