package roller

import "testing"

func TestRoller(t *testing.T) {
	// good inputs
	for _, tt := range []struct {
		name, input string
	}{
		{
			name:  "d4 success",
			input: "d4",
		},
		{
			name:  "d6 success",
			input: "d6",
		},
		{
			name:  "d8 success",
			input: "d8",
		},
		{
			name:  "d10 success",
			input: "d10",
		},
		{
			name:  "d12 success",
			input: "d12",
		},
		{
			name:  "d20 success",
			input: "d20",
		},
		{
			name:  "d100 success",
			input: "d100",
		},
		{
			name:  "multiple dice success",
			input: "2d6",
		},
		{
			name:  "dice with modifier success",
			input: "3d4+1",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Roll(tt.input)
			if err != nil {
				t.Errorf("unexpected error for input %s: %s", tt.input, err)
			}
		})
	}

	// bad inputs
	for _, tt := range []struct {
		name, input string
	}{
		{
			name:  "less than 1 dice failure",
			input: "0d6",
		},
		{
			name:  "less than 1 side failure",
			input: "2d0",
		},
		{
			name:  "more than 10,000 dice failure",
			input: "10001d6",
		},
		{
			name:  "non-standard dice size failure",
			input: "d17",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Roll(tt.input)
			if err == nil {
				t.Errorf("expected error for input %s but did not get one", tt.input)
			}
		})
	}
}

func FuzzRoller(f *testing.F) {
	f.Add("2d6")
	f.Add("d20")
	f.Add("1d4+1")
	f.Fuzz(func(t *testing.T, input string) {
		_, err := Roll(input)
		if err != nil {
			return
		}
	})
}
