package patterns

import (
	"math/rand"
)

type (
	BoolGeneratorSlow struct {
		chance int
	}
)

// NewBoolGeneratorSlow returns a new BoolGeneratorSlow with the given chance.
// 1 = 100%
// 2 = 50%
// 3 = 33.3%
// 4 = 25%
// 5 = 20%
// 6 = 16.6%
// 7 = 14.2%
// 8 = 12.5%
// 9 = 11.1%
// 10 = 10%
// etc
func NewBoolGeneratorSlow(chance int) *BoolGeneratorSlow {
	return &BoolGeneratorSlow{
		chance: chance,
	}
}

func (bgs *BoolGeneratorSlow) Bool() bool {
	return bgs.chance == rand.Intn(bgs.chance)+1
}
