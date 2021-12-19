package roulette

type Roulette interface {
	Name() string
	NumberCount() int
	PayoutWith(result int, s Strategy) int
}
