package unittest

import "testing"

func TestHomeworkBE4905186(t *testing.T) {
	t.Run("should return false if characters less than 3", func(t *testing.T) {
		input := "ab"

		want := false

		got, _ := HomeworkBE4905186(input)
		if got != want {
			t.Errorf("HomeworkBE4905186(%v) = false, want %v", got, want)
		}
	})

	t.Run("should return false if characters more than 15", func(t *testing.T) {
		input := "abcdefghijklmnopqrstuvwqyz"

		want := false

		got, _ := HomeworkBE4905186(input)
		if got != want {
			t.Errorf("HomeworkBE4905186(%v) = false, want %v", got, want)
		}
	})

	t.Run("should return false if not palindrome", func(t *testing.T) {
		input := "race"

		want := false

		got, _ := HomeworkBE4905186(input)
		if got != want {
			t.Errorf("HomeworkBE4905186(%v) = false, want %v", got, want)
		}
	})

	t.Run("should return true if palindrome", func(t *testing.T) {
		input := "racecar"

		want := true

		got, _ := HomeworkBE4905186(input)
		if got != want {
			t.Errorf("HomeworkBE4905186(%v) = false, want %v", got, want)
		}
	})
}
