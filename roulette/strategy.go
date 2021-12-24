package roulette

type BetType uint8

const (
	Unknown BetType = iota
	Zero
	DoubleZero
	StraightUp
	Row
	Split
	Street
	Corner
	BasketEuropean
	BasketUS
	DoubleStreet
	FirstColumn
	SecondColumn
	ThirdColumn
	FirstDozen
	SecondDozen
	ThirdDozen
	Odd
	Even
	Red
	Black
	FirstHalf
	SecondHalf
	end
)

var (
	NumbersRed   []int = []int{32, 19, 21, 25, 34, 27, 36, 30, 23, 5, 16, 1, 14, 9, 18, 7, 12, 3}
	NumbersBlack []int = []int{15, 4, 2, 17, 6, 13, 11, 8, 10, 24, 33, 20, 31, 22, 29, 28, 35, 26}
)

type Bet struct {
	Amount  int
	BetType BetType
	Numbers []int
}

func NewBet(amount int, betType BetType, numbers ...int) Bet {
	return Bet{
		Amount:  amount,
		BetType: betType,
		Numbers: numbers,
	}
}

type Strategy struct {
	name string
	bets []Bet
}

func NewStrategy(name string) Strategy {
	return Strategy{
		name: name,
	}
}

func (s *Strategy) AddBet(bet Bet) {
	s.bets = append(s.bets, bet)
}

func (s Strategy) MinimalBet() int {
	var minimalAmount int
	for _, bet := range s.bets {
		minimalAmount += bet.Amount
	}
	return minimalAmount
}

func StrategyOdd() Strategy {
	s := NewStrategy("Odd")
	s.AddBet(NewBet(1, Odd))
	return s
}

func StrategyEven() Strategy {
	s := NewStrategy("Even")
	s.AddBet(NewBet(1, Even))
	return s
}

func StrategyRed() Strategy {
	s := NewStrategy("Red")
	s.AddBet(NewBet(1, Red, NumbersRed...))
	return s
}

func StrategyBlack() Strategy {
	s := NewStrategy("Black")
	s.AddBet(NewBet(1, Black, NumbersBlack...))
	return s
}

func StrategyBiColor() Strategy {
	s := NewStrategy("Bi color")
	s.AddBet(NewBet(1, Black, NumbersBlack...))
	s.AddBet(NewBet(1, Red, NumbersRed...))
	return s
}

func StrategyDoubleStreetQuad() Strategy {
	s := NewStrategy("Double Street Quad")
	s.AddBet(NewBet(1, StraightUp, 1))
	s.AddBet(NewBet(1, StraightUp, 0))
	s.AddBet(NewBet(2, DoubleStreet, 1, 2, 3, 4, 5, 6))
	s.AddBet(NewBet(2, DoubleStreet, 7, 8, 9, 10, 11, 12))
	s.AddBet(NewBet(1, Corner, 13, 14, 16, 17))
	return s
}

func StrategyZero() Strategy {
	s := NewStrategy("Zero only")
	s.AddBet(NewBet(1, StraightUp, 0))
	return s
}
