package unittest

import "testing"

func TestHomeworkBE5132238(t *testing.T) {
	tests := []struct {
		name   string
		number int
		want   int
	}{
		{
			name:   "number is greater than 8",
			number: 9,
			want:   -1,
		},
		{
			name:   "number is less than 0",
			number: -1,
			want:   0,
		},
		{
			name:   "number is 0",
			number: 0,
			want:   1,
		},
		{
			name:   "number is between 1 and 8",
			number: 5,
			want:   120,
			// 5! = 5 * 4 * 3 * 2 * 1 = 120
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HomeworkBE5132238(tt.number); got != tt.want {
				t.Errorf("HomeworkBE5132238(%d) = %v, want %v", tt.number, got, tt.want)
			}
		})
	}
}
