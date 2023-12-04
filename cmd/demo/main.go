package main

import (
	"math/rand"
	"os"
	"roulette/entity"
	"roulette/usecase/entropy"
	"roulette/usecase/roulette"
	"roulette/usecase/strategy"
	"time"
)

// https://beechplane.wordpress.com/2012/03/06/visualizing-probability-roulette/

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	r := roulette.NewFrench()
	conf := entity.SimulationConf{
		Writer:      os.Stdout,
		Entropy:     entropy.NewRandom(r.NumberCount()),
		Roulette:    r,
		NbRun:       100000,
		MaxSpins:    100,
		StartAmount: 100,
	}
	sim := entity.NewSimulation(conf)
	sim.RunWith(strategy.BiColor())
	sim.RunWith(strategy.Odd())
	sim.RunWith(strategy.Even())
	sim.RunWith(strategy.Red())
	sim.RunWith(strategy.Black())
	sim.RunWith(strategy.DoubleStreetQuad())
	sim.RunWith(strategy.Zero())
	sim.RunWith(strategy.CompleteBet())
}
