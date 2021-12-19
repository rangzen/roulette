package roulette

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRouletteFrenchOddWin(t *testing.T) {
	f := NewFrenchRoulette()
	s := NewStrategy("Odd")
	s.AddBet(NewBet(1, Odd))

	payout := f.PayoutWith(1, s)

	assert.Equal(t, 2, payout)
}

func TestRouletteFrenchOddLost(t *testing.T) {
	f := NewFrenchRoulette()
	s := NewStrategy("Odd")
	s.AddBet(NewBet(1, Odd))

	payout := f.PayoutWith(2, s)

	assert.Equal(t, -1, payout)
}

func TestRouletteFrenchEvenWin(t *testing.T) {
	f := NewFrenchRoulette()
	s := NewStrategy("Even")
	s.AddBet(NewBet(1, Even))

	payout := f.PayoutWith(2, s)

	assert.Equal(t, 2, payout)
}

func TestRouletteFrenchEvenLost(t *testing.T) {
	f := NewFrenchRoulette()
	s := NewStrategy("Even")
	s.AddBet(NewBet(1, Even))

	payout := f.PayoutWith(1, s)

	assert.Equal(t, -1, payout)
}
