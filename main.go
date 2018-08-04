package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/valakuzhyk/planetdomains/game"
)

func main() {
	game.StartupMenu()
	p1, p2, err := game.CreateFor2()
	if err != nil {
		log.Fatal("Unable to create Game:", err)
	}
	for true {
		fmt.Printf("%+v\n", p1)
		p1.PlayHand(p2)
		fmt.Printf("%+v\n", p1)
		p2.PlayHand(p1)
	}
}
