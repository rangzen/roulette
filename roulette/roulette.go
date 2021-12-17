package roulette

type Roulette interface {
	SpinOn(s Strategy) int
}
