package roller

import "testing"

func TestRoller(t *testing.T) {
	t.Run("rolls successfully", func(t *testing.T) {
		num, sides, modifier := 2, 6, 0
		min := num*1 + modifier
		max := num*sides + modifier

		got, _ := Roll(num, sides, modifier)
		if got < min || got > max {
			t.Errorf("result outside of range: %d", got)
		}
	})

	t.Run("errors if less than 1 dice", func(t *testing.T) {
		num, sides, modifier := 0, 6, 0
		_, err := Roll(num, sides, modifier)
		if err == nil {
			t.Error("expected error")
		}
	})

	t.Run("errors if less than 1 side", func(t *testing.T) {
		num, sides, modifier := 2, 0, 0
		_, err := Roll(num, sides, modifier)
		if err == nil {
			t.Error("expected error")
		}
	})
}

func FuzzRoller(f *testing.F) {
	f.Add(2, 6, 0)
	f.Fuzz(func(t *testing.T, num, sides, mod int) {
		min := num*1 + mod
		max := num*sides + mod

		got, err := Roll(num, sides, mod)
		if err != nil {
			return
		}
		if got < min || got > max {
			t.Errorf("result outside of range: %d", got)
		}
	})
}
