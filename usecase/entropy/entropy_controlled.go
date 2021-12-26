package entropy

import "roulette/entity"

type Controlled struct {
	target int
}

func NewControlled(target int) entity.Entropy {
	return Controlled{
		target: target,
	}
}

func (c Controlled) Spin() int {
	return c.target
}
