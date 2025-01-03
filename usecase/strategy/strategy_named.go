package strategy

import "roulette/entity"

func Odd() entity.Strategy {
	s := entity.NewStrategy("Odd")
	s.AddBet(entity.NewBet(1, entity.Odd))
	return s
}

func Even() entity.Strategy {
	s := entity.NewStrategy("Even")
	s.AddBet(entity.NewBet(1, entity.Even))
	return s
}

func Red() entity.Strategy {
	s := entity.NewStrategy("Red")
	s.AddBet(entity.NewBet(1, entity.Red, entity.NumbersRed...))
	return s
}

func Black() entity.Strategy {
	s := entity.NewStrategy("Black")
	s.AddBet(entity.NewBet(1, entity.Black, entity.NumbersBlack...))
	return s
}

func BiColor() entity.Strategy {
	s := entity.NewStrategy("Bi color")
	s.AddBet(entity.NewBet(1, entity.Black, entity.NumbersBlack...))
	s.AddBet(entity.NewBet(1, entity.Red, entity.NumbersRed...))
	return s
}

func DoubleStreetQuad() entity.Strategy {
	s := entity.NewStrategy("Double Street Quad")
	s.AddBet(entity.NewBet(1, entity.StraightUp, 1))
	s.AddBet(entity.NewBet(1, entity.StraightUp, 0))
	s.AddBet(entity.NewBet(2, entity.DoubleStreet, 1, 2, 3, 4, 5, 6))
	s.AddBet(entity.NewBet(2, entity.DoubleStreet, 7, 8, 9, 10, 11, 12))
	s.AddBet(entity.NewBet(1, entity.Corner, 13, 14, 16, 17))
	return s
}

func Zero() entity.Strategy {
	s := entity.NewStrategy("Zero only")
	s.AddBet(entity.NewBet(1, entity.StraightUp, 0))
	return s
}

func CompleteBet() entity.Strategy {
	s := entity.NewStrategy("Full bet (17)")
	s.AddBet(entity.NewBet(1, entity.StraightUp, 17))
	s.AddBet(entity.NewBet(1, entity.Split, 17, 14))
	s.AddBet(entity.NewBet(1, entity.Split, 17, 18))
	s.AddBet(entity.NewBet(1, entity.Split, 17, 16))
	s.AddBet(entity.NewBet(1, entity.Split, 17, 20))
	s.AddBet(entity.NewBet(1, entity.Corner, 17, 14, 15, 18))
	s.AddBet(entity.NewBet(1, entity.Corner, 17, 18, 20, 21))
	s.AddBet(entity.NewBet(1, entity.Corner, 17, 20, 19, 16))
	s.AddBet(entity.NewBet(1, entity.Corner, 17, 16, 13, 14))
	return s
}

func ThreeTwoBet() entity.Strategy {
	s := entity.NewStrategy("3/2 Black version bet")
	s.AddBet(entity.NewBet(3, entity.Black, entity.NumbersBlack...))
	s.AddBet(entity.NewBet(2, entity.ThirdColumn, entity.NumbersThirdColumn...))
	return s
}
