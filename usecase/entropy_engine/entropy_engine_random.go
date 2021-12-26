package entropy_engine

import (
	"math/rand"
	"roulette/entity"
)

type Random struct {
	interval int
}

func NewRandom(nb int) entity.EntropyEngine {
	return Random{
		interval: nb,
	}
}

func (r Random) Spin() int {
	return int(rand.Int31n(int32(r.interval)))
}
