package unittest

import (
    "testing"
)

func TestHomeworkBE5188276(t *testing.T) {
    tests := []struct {
        input int
        expected bool
    }{
        {2, true},
        {5, true},
        {11, true},
        {15, false},
        {17, true},
        {20, false},
        {23, true},
		{25, false},
        {29, true},
        {30, false},
        {32, false},
        {43, true},
        {46, false},
        {51, false},
    }

    for _, test := range tests {
        result := HomeworkBE5188276(test.input)
        if result != test.expected {
            t.Errorf("HomeworkBE5188276(%d) = %t; expected %t", test.input, result, test.expected)
        }
    }

    result := HomeworkBE5188276(51)
    if result != false {
        t.Errorf("HomeworkBE5188276(51) = %t; expected false", result)
    }

    result = HomeworkBE5188276(0)
    if result != false {
        t.Errorf("HomeworkBE5188276(0) = %t; expected false", result)
    }
}

