package roulette

import (
	"math/rand"
)

type RouletteFrench struct{}

func NewFrenchRoulette() Roulette {
	return RouletteFrench{}
}

func (f RouletteFrench) SpinOn(s Strategy) int {
	spin := int(rand.Int31n(37))

	var totalWin int
	for _, b := range s.bets {
		var payout = -1
		if b.BetType == Odd && spin%2 == 1 {
			payout = 1
		} else if b.BetType == Even && spin%2 == 0 {
			payout = 1
		} else {
			for _, n := range b.Numbers {
				if spin == n {
					switch b.BetType {
					case StraightUp:
						payout = 35
					case Corner:
						payout = 8
					case DoubleStreet:
						payout = 5
					case Red:
						payout = 1
					case Black:
						payout = 1
					default:
						panic("unknown bet type")
					}
					break
				}
			}
		}
		totalWin += b.Amount * payout
	}
	return totalWin
}
