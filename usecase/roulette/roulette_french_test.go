package roulette_test

import (
	"github.com/stretchr/testify/assert"
	"roulette/entity"
	"roulette/usecase/roulette"
	"testing"
)

func TestRouletteFrenchEvenLooseWith0(t *testing.T) {
	r := roulette.French{}
	s := entity.NewStrategy("Even")
	s.AddBet(entity.NewBet(1, entity.Even))

	p := r.PayoutWith(0, s)

	assert.Equal(t, -1, p)
}

func TestRouletteFrenchOddWin(t *testing.T) {
	r := roulette.French{}
	s := entity.NewStrategy("Odd")
	s.AddBet(entity.NewBet(1, entity.Odd))

	p := r.PayoutWith(1, s)

	assert.Equal(t, 2, p)
}

func TestRouletteFrenchOddLost(t *testing.T) {
	r := roulette.French{}
	s := entity.NewStrategy("Odd")
	s.AddBet(entity.NewBet(1, entity.Odd))

	p := r.PayoutWith(2, s)

	assert.Equal(t, -1, p)
}

func TestRouletteFrenchEvenWin(t *testing.T) {
	r := roulette.French{}
	s := entity.NewStrategy("Even")
	s.AddBet(entity.NewBet(1, entity.Even))

	p := r.PayoutWith(2, s)

	assert.Equal(t, 2, p)
}

func TestRouletteFrenchEvenLost(t *testing.T) {
	r := roulette.French{}
	s := entity.NewStrategy("Even")
	s.AddBet(entity.NewBet(1, entity.Even))

	p := r.PayoutWith(1, s)

	assert.Equal(t, -1, p)
}
