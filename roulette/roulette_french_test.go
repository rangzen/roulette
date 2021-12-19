package roulette

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRouletteFrenchOddWin(t *testing.T) {
	f := NewFrenchRoulette(NewControlledEngine(1))
	s := NewStrategy("Odd")
	s.AddBet(NewBet(1, Odd))

	payout := f.SpinOn(s)

	assert.Equal(t, 1, payout)
}

func TestRouletteFrenchOddLost(t *testing.T) {
	f := NewFrenchRoulette(NewControlledEngine(2))
	s := NewStrategy("Odd")
	s.AddBet(NewBet(1, Odd))

	payout := f.SpinOn(s)

	assert.Equal(t, -1, payout)
}

func TestRouletteFrenchEvenWin(t *testing.T) {
	f := NewFrenchRoulette(NewControlledEngine(2))
	s := NewStrategy("Even")
	s.AddBet(NewBet(1, Even))

	payout := f.SpinOn(s)

	assert.Equal(t, 1, payout)
}

func TestRouletteFrenchEvenLost(t *testing.T) {
	f := NewFrenchRoulette(NewControlledEngine(1))
	s := NewStrategy("Even")
	s.AddBet(NewBet(1, Even))

	payout := f.SpinOn(s)

	assert.Equal(t, -1, payout)
}
