package roller

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func parse(input string) (num, sides, mod int, err error) {
	re := regexp.MustCompile(`^(\d+)?d(\d+)([+-]\d+)?$`)
	matches := re.FindStringSubmatch(input)
	if matches == nil {
		err = errors.New("invalid input string")
		return
	}

	if matches[1] != "" {
		var num64 int64
		num64, err = strconv.ParseInt(matches[1], 0, 0)
		if err != nil {
			return
		}
		num = int(num64)
	} else {
		num = 1
	}
	if num < 1 {
		err = fmt.Errorf("invalid number of dice in %s: %d", input, num)
		return
	}

	sides64, err := strconv.ParseInt(matches[2], 0, 0)
	if err != nil {
		return
	}
	sides = int(sides64)
	if sides < 1 {
		err = fmt.Errorf("invalid number of sides in %s: %d", input, sides)
		return
	}

	if matches[3] != "" {
		var mod64 int64
		mod64, err = strconv.ParseInt(matches[3], 0, 0)
		if err != nil {
			return
		}
		mod = int(mod64)
	}

	return
}
