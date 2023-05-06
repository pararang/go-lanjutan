package unittest

import "testing"

func TestHomeworkBE5011795(t *testing.T) {
	t.Run("should return false if characters less than 3", func(t *testing.T) {
		input := "al"

		want := false

		got, _ := HomeworkBE5011795(input)
		if got != want {
			t.Errorf("HomeworkBE5011795(%v) = false, want %v", got, want)
		}
	})

	t.Run("should return false if characters more than 15", func(t *testing.T) {
		input := "albitegarprayoga"

		want := false

		got, _ := HomeworkBE5011795(input)
		if got != want {
			t.Errorf("HomeworkBE5011795(%v) = false, want %v", got, want)
		}
	})

	t.Run("should return false if not palindrome", func(t *testing.T) {
		input := "kode"

		want := false

		got, _ := HomeworkBE5011795(input)
		if got != want {
			t.Errorf("HomeworkBE5011795(%v) = false, want %v", got, want)
		}
	})

	t.Run("should return true if palindrome", func(t *testing.T) {
		input := "rececer"

		want := true

		got, _ := HomeworkBE5011795(input)
		if got != want {
			t.Errorf("HomeworkBE5011795(%v) = false, want %v", got, want)
		}
	})
}
