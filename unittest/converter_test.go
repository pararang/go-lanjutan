package unittest

import "testing"

func TestWide_ToMeter(t *testing.T) {
	t.Run("should return existing value if unit already m", func(t *testing.T) {
		w := Wide{
			Unit:  "m",
			Value: 4.5,
		}

		want := 4.5

		got := w.ToMeter()
		if got != want {
			t.Errorf("Wide.ToMeter() = %v, want %v", got, want)
		}
	})

	t.Run("should return -1 if unit is empty", func(t *testing.T) {
		w := Wide{
			Value: 4.5,
		}

		want := float64(-1)

		got := w.ToMeter()
		if got != want {
			t.Errorf("Wide.ToMeter() = %v, want %v", got, want)
		}
	})

	// t.Run("should convert from cm correctly", func(t *testing.T) {
	// 	w := Wide{
	// 		Value: 200,
	// 		Unit:  "cm",
	// 	}

	// 	want := 2.0

	// 	got := w.ToMeter()
	// 	if got != want {
	// 		t.Errorf("Wide.ToMeter() = %v, want %v", got, want)
	// 	}
	// })
}
