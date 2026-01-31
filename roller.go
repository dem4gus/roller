package roller

import (
	"math/rand/v2"
)

type DiceSet struct {
	num   int
	sides int
	mod   int
}

func Roll(input string) (int, error) {
	d, err := NewDiceSet(input)
	if err != nil {
		return 0, err
	}
	return d.Roll(), nil
}

func NewDiceSet(input string) (*DiceSet, error) {
	num, sides, mod, err := parse(input)
	if err != nil {
		return nil, err
	}

	return &DiceSet{num, sides, mod}, nil
}

func (d DiceSet) Roll() int {
	result := d.mod
	for range d.num {
		result += rand.N(d.sides) + 1
	}
	return result
}
