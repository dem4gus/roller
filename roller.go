package roller

import (
	"errors"
	"math/rand/v2"
)

func Roll(num, sides, modifier int) (int, error) {
	if num < 1 {
		return 0, errors.New("cannot roll less than one dice")
	}
	if sides < 1 {
		return 0, errors.New("dice cannot have less than one side")
	}

	result := 0
	for range num {
		result += rand.N(sides) + 1
	}
	return result + modifier, nil
}
