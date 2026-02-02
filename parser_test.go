package roller

import "testing"

func TestParse(t *testing.T) {
	for _, tt := range []struct {
		name  string
		input string
		num   int
		sides int
		mod   int
	}{
		{
			name:  "explicit single die",
			input: "1d6",
			num:   1,
			sides: 6,
		},
		{
			name:  "implied single die",
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
	} {
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
}

func FuzzParse(f *testing.F) {
	for _, input := range []string{"1d6", "d6", "2d6", "2d6+2", "2d6-2"} {
		f.Add(input)
	}
	f.Fuzz(func(t *testing.T, input string) {
		num, sides, _, err := parse(input)
		if num < 0 {
			t.Errorf("num should not be negative")
		}
		if sides < 0 {
			t.Errorf("sides should not be negative")
		}
		if err != nil {
			return
		}
	})
}
