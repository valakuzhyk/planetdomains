package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/valakuzhyk/planetdomains/game"
)

func main() {
	// game.StartupMenu()
	field, err := game.CreateFor2()
	if err != nil {
		log.Fatal("Unable to create Game:", err)
	}
	for true {
		field.PlayHand()
	}
}
