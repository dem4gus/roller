package roller

import "testing"

func TestRoller(t *testing.T) {
	t.Run("rolls successfully", func(t *testing.T) {
		input := "2d6"

		_, err := Roll(input)
		if err != nil {
			t.Errorf("unexpected error %s", err)
		}
	})

	t.Run("errors if less than 1 dice", func(t *testing.T) {
		input := "0d6"
		_, err := Roll(input)
		if err == nil {
			t.Errorf("expected error but did not get one")
		}
	})

	t.Run("errors if less than 1 side", func(t *testing.T) {
		input := "2d0"
		_, err := Roll(input)
		if err == nil {
			t.Error("expected error but did not get one")
		}
	})

	t.Run("only allows 9,999 dice maximum", func(t *testing.T) {
		input := "10000d6"
		_, err := Roll(input)
		if err == nil {
			t.Error("expected error but did not get one")
		}
	})
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
