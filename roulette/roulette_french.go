package roulette

type RouletteFrench struct {
	randomEngine EntropyEngine
}

func NewFrenchRoulette(r EntropyEngine) Roulette {
	return RouletteFrench{
		randomEngine: r,
	}
}

func (f RouletteFrench) Spin() int {
	return f.randomEngine.Spin()
}

func (f RouletteFrench) PayoutWith(result int, s Strategy) int {
	var totalWin int
	for _, b := range s.bets {
		var payout = -1
		if b.BetType == Odd && result%2 == 1 {
			payout = 1 + 1
		} else if b.BetType == Even && result%2 == 0 {
			payout = 1 + 1
		} else {
			for _, n := range b.Numbers {
				if result == n {
					switch b.BetType {
					case StraightUp:
						payout = 1 + 35
					case Corner:
						payout = 1 + 8
					case DoubleStreet:
						payout = 1 + 5
					case Red:
						payout = 1 + 1
					case Black:
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
