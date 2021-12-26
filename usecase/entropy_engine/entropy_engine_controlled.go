package entropy_engine

import "roulette/entity"

type Controlled struct {
	target int
}

func NewControlled(target int) entity.EntropyEngine {
	return Controlled{
		target: target,
	}
}

func (c Controlled) Spin() int {
	return c.target
}
