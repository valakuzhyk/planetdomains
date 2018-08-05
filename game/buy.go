package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/valakuzhyk/planetdomains/card"
	"github.com/valakuzhyk/planetdomains/utils"

	"github.com/buger/goterm"
)

func purchaseCards(p *person) {
	// If possible try and purchase some cards
	// 0 represents explorer
	for true {
		goterm.Clear()
		goterm.Flush()
		choice := getPurchaseChoice(p.Trade, p.field)

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

func getPurchaseChoice(playerTrade int, field *Field) int {
	cardsToSelectFrom := field.TradeRow.Cards
	if !field.Explorers.IsEmpty() {
		cardsToSelectFrom = append([]card.Card{field.Explorers.Peek()}, cardsToSelectFrom...)
	}

	i := utils.PickCard(
		fmt.Sprintf("You have %d trade, what would you like to buy?", playerTrade),
		cardsToSelectFrom,
		false /* not required */)
	return i
}

func buyExplorer(p *person) {
	explorers := p.field.Explorers
	if explorers.IsEmpty() {
		fmt.Println("No more explorers!")
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
	if n > len(tradeRow.Cards) {
		fmt.Println("Invalid choice\n")
		return
	}

	// User input is 1-indexed
	c := tradeRow.Peek(n - 1)
	if p.GetTrade() < c.GetCost() {
		fmt.Printf("Cost %d greater than your trade %d. Insufficient Funds\n",
			c.GetCost(),
			p.GetTrade())
		return
	}

	c = p.field.TradeRow.Take(n - 1)
	p.AddTrade(-c.GetCost())

	// Replace the card on the trade row if possible
	if !p.field.TradeDeck.IsEmpty() {
		p.field.TradeRow.Insert(n-1, p.field.TradeDeck.Draw())
	}

	p.Discard.PlaceOnTop(c)
}
