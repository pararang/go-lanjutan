package unittest

import (
	"testing"
)

func TestHomeworkBE4290472(t *testing.T) {
	for _, v := range []int{1, 10, 20, 30, 40, 51} {
		if HomeworkBE4290472(v) != false {
			t.Errorf("HomeworkBE4290472(%v) = false; expected true", v)
		}
	}

	for _, v := range []int{11, 29, 31, 41} {
		if HomeworkBE4290472(v) != true {
			t.Errorf("HomeworkBE4290472(%v) = true; expected false", v)
		}
	}
}
