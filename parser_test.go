package roller

import "testing"

func TestParser(t *testing.T) {
	happyPath := []struct {
		name  string
		input string
		num   int
		sides int
		mod   int
	}{
		{
			name:  "single die",
			input: "d6",
			num:   1,
			sides: 6,
		},
		{
			name:  "multiple dice, no modifier",
			input: "2d6",
			num:   2,
			sides: 6,
		},
		{
			name:  "with positive modifier",
			input: "2d6+2",
			num:   2,
			sides: 6,
			mod:   2,
		},
		{
			name:  "with negative modifier",
			input: "2d6-2",
			num:   2,
			sides: 6,
			mod:   -2,
		},
		{
			name:  "max number of dice",
			input: "9999d6",
			num:   9999,
			sides: 6,
		},
	}

	for _, tt := range happyPath {
		t.Run(tt.name, func(t *testing.T) {
			num, sides, mod, err := parse(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if tt.num != num {
				t.Errorf("wanted %d dice but got %d", tt.num, num)
			}
			if tt.sides != sides {
				t.Errorf("wanted %d sides but got %d", tt.sides, sides)
			}
			if tt.mod != mod {
				t.Errorf("wanted modifier %d but got %d", tt.mod, mod)
			}
		})
	}

	errorCases := []struct {
		name  string
		input string
	}{
		{
			name:  "too many dice",
			input: "10000d6",
		},
	}

	for _, tt := range errorCases {
		t.Run(tt.name, func(t *testing.T) {
			_, _, _, err := parse(tt.input)
			if err == nil {
				t.Fatalf("expected error but got none")
			}
		})
	}
}
