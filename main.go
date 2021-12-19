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
	r := roulette.NewFrenchRoulette()
	simConf := roulette.SimulationConf{
		Writer:        os.Stdout,
		EntropyEngine: roulette.NewRandomEngine(r.NumberCount()),
		Roulette:      r,
		NbRun:         100000,
		MaxSpins:      100,
		StartAmount:   100,
	}
	s := roulette.NewSimulation(simConf)
	s.RunWith(roulette.StrategyBiColor())
	s.RunWith(roulette.StrategyOdd())
	s.RunWith(roulette.StrategyRed())
	s.RunWith(roulette.StrategyDoubleStreetQuad())
	s.RunWith(roulette.StrategyZero())
}
