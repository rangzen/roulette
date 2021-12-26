package entity

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
	Name       string
	Bets       []Bet
	minimalBet int
}

func NewStrategy(name string) Strategy {
	return Strategy{
		Name: name,
	}
}

func (s *Strategy) AddBet(bet Bet) {
	s.Bets = append(s.Bets, bet)

	var minimalAmount int
	for _, b := range s.Bets {
		minimalAmount += b.Amount
	}
	s.minimalBet = minimalAmount
}

func (s Strategy) MinimalBet() int {
	return s.minimalBet
}
