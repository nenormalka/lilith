package patterns

import (
	"math/rand"
	"sync"
	"time"
)

const (
	ChanceHalf    Chance = "50%"
	ChanceQuarter        = "25%"
	ChanceEighth         = "12.5%"
)

type (
	BoolGenerator struct {
		src       rand.Source
		mu        sync.Mutex
		cache     int64
		rate      int64
		remaining int
	}

	Chance string
)

func NewBoolGenerator(c Chance) *BoolGenerator {
	bg := &BoolGenerator{
		src: rand.NewSource(time.Now().UnixNano()),
		mu:  sync.Mutex{},
	}

	bg.setRate(c)

	return bg
}

func (bg *BoolGenerator) Bool() bool {
	bg.mu.Lock()
	defer bg.mu.Unlock()

	if bg.remaining == 0 {
		bg.cache, bg.remaining = bg.src.Int63(), 63
	}

	result := bg.cache&bg.rate == bg.rate
	bg.cache >>= 1
	bg.remaining--

	return result
}

func (bg *BoolGenerator) setRate(chance Chance) {
	switch chance {
	case ChanceHalf:
		bg.rate = 2
	case ChanceQuarter:
		bg.rate = 3
	case ChanceEighth:
		bg.rate = 7
	default:
		bg.rate = 2
	}
}
