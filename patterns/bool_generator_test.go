package patterns

import (
	"fmt"
	"os"
	"testing"
)

func TestBoolGenerator(t *testing.T) {
	bg := NewBoolGenerator(ChanceEighth)

	fo, err := os.Create("chance.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = fo.Close(); err != nil {
			panic(err)
		}
	}()

	for i := 0; i < 10000; i++ {
		if _, err = fo.Write([]byte(fmt.Sprintf("%v \n", bg.Bool()))); err != nil {
			panic(err)
		}
	}
}
