package roller

import (
	"errors"
	"regexp"
	"strconv"
)

func parse(input string) (num, sides, mod int, err error) {
	re := regexp.MustCompile(`^(\d+)?d(\d+)([+-]\d+)?$`)
	// matches[1]: number of dice
	// matches[2]: number of sides
	// matches[3]: modifier
	matches := re.FindStringSubmatch(input)
	if matches == nil {
		err = errors.New("invalid input string")
		return num, sides, mod, err
	}

	if matches[1] == "" {
		num = 1
	} else {
		var num64 int64
		num64, err = strconv.ParseInt(matches[1], 0, 0)
		if err != nil {
			return num, sides, mod, err
		}
		num = int(num64)
	}

	sides64, err := strconv.ParseInt(matches[2], 0, 0)
	if err != nil {
		return num, sides, mod, err
	}
	sides = int(sides64)

	if matches[3] != "" {
		var mod64 int64
		mod64, err = strconv.ParseInt(matches[3], 0, 0)
		if err != nil {
			return num, sides, mod, err
		}
		mod = int(mod64)
	}

	return num, sides, mod, err
}
