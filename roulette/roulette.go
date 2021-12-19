package roulette

import "math/rand"

type Roulette interface {
	Spin() int
	PayoutWith(result int, s Strategy) int
}

type EntropyEngine interface {
	Spin() int
}

type RandomEngine struct {
	interval int
}

func NewRandomEngine(nb int) EntropyEngine {
	return RandomEngine{
		interval: nb,
	}
}

func (r RandomEngine) Spin() int {
	return int(rand.Int31n(int32(r.interval)))
}

type ControlledEngine struct {
	target int
}

func NewControlledEngine(target int) EntropyEngine {
	return ControlledEngine{
		target: target,
	}
}

func (c ControlledEngine) Spin() int {
	return c.target
}
