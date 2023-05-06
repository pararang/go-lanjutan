package unittest

import (
	"testing"
)

func TestHomeworkBE5227642(t *testing.T) {
	test_1 := "" // 0 char

	if HomeworkBE5227642(test_1) != "add more chars" {
		t.Errorf("HomeworkBE4290472(%v) expected 'add more chars'", test_1)
	}

	test_2 := "abcdefghijklmnopqrstuvwxyz" // 25 chars

	if HomeworkBE5227642(test_2) != "to much" {
		t.Errorf("HomeworkBE4290472(%v) expected 'to much'", test_2)
	}

	test_3 := "1234567890" // 10 chars

	if HomeworkBE5227642(test_3) != "0987654321" {
		t.Errorf("HomeworkBE4290472(%v) expected '0987654321'", test_3)
	}
}
