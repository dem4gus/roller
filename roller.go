package roller

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

const (
	minDice = 1
	maxDice = 10_000
)

var diceSizes = []int{4, 6, 8, 10, 12, 20, 100}

func Roll(input string) (int, error) {
	d, err := NewDiceSet(input)
	if err != nil {
		return 0, err
	}
	return d.Roll(), nil
}

type DiceSet struct {
	num   int
	sides int
	mod   int
}

func NewDiceSet(input string) (*DiceSet, error) {
	num, sides, mod, err := parse(input)
	if err != nil {
		return nil, err
	}
	if num < minDice || num > maxDice {
		return nil, fmt.Errorf("invalid number of dice in %s: %d", input, num)
	}
	if !slices.Contains(diceSizes, sides) {
		return nil, fmt.Errorf("invalid number of sides in %s: %d", input, sides)
	}

	return &DiceSet{num, sides, mod}, nil
}

func (d DiceSet) Roll() int {
	result := d.mod
	for range d.num {
		result += rand.N(d.sides) + 1 // #nosec G404 -- not cryptographic
	}
	return result
}
