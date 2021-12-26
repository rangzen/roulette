package roulette

import "roulette/entity"

type French struct {
}

func NewFrench() entity.Roulette {
	return French{}
}

func (f French) Name() string {
	return "French Roulette"
}

func (f French) NumberCount() int {
	return 37
}

func (f French) PayoutWith(result int, s entity.Strategy) int {
	var totalWin int
	for _, b := range s.Bets {
		var payout = -1
		if b.BetType == entity.Odd && result%2 == 1 {
			payout = 1 + 1
		} else if b.BetType == entity.Even && result >= 2 && result%2 == 0 {
			payout = 1 + 1
		} else {
			for _, n := range b.Numbers {
				if result == n {
					switch b.BetType {
					case entity.StraightUp:
						payout = 1 + 35
					case entity.Corner:
						payout = 1 + 8
					case entity.DoubleStreet:
						payout = 1 + 5
					case entity.Red:
						payout = 1 + 1
					case entity.Black:
						payout = 1 + 1
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
