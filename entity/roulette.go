package entity

// Roulette describes a type of roulette, e.g., French, American, Triple-zero wheel, etc.
type Roulette interface {
	Name() string
	NumberCount() int
	PayoutWith(result int, s Strategy) int
}
