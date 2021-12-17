package main

import (
	"math/rand"
	"os"
	"roulette/roulette"
	"time"
)

// https://beechplane.wordpress.com/2012/03/06/visualizing-probability-roulette/

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	simConf := roulette.SimulationConf{
		Writer:      os.Stdout,
		Roulette:    roulette.NewFrenchRoulette(),
		NbRun:       100000,
		MaxSpins:    250,
		StartAmount: 10,
	}
	s := roulette.NewSimulation(simConf)
	s.RunWith(roulette.StrategyBiColor())
	s.RunWith(roulette.StrategyOdd())
	s.RunWith(roulette.StrategyRed())
	s.RunWith(roulette.StrategyDoubleStreetQuad())
}
