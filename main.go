package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/valakuzhyk/planetdomains/game"
)

func main() {
	p1, p2, err := game.CreateFor2()
	if err != nil {
		log.Fatal("Unable to create Game:", err)
	}
	for true {
		log.Infof("%+v", p1)
		p1.PlayHand(p2)
		log.Infof("%+v", p1)
		p2.PlayHand(p1)
	}
}
