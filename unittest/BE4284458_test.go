package unittest

import "testing"

func TestHomeworkBE4284458(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "Numbers more than 25",
			args: args{numbers: []int{25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
			want: -1,
		},
		{
			name: "Largest number",
			args: args{numbers: []int{1, 9, 4, 2}},
			want: 9,
		},
		{
			name: "Largest number",
			args: args{numbers: []int{1, 9, 0, 2}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HomeworkBE4284458(tt.args.numbers); got != tt.want {
				t.Errorf("HomeworkBE4284458() = %v, want %v", got, tt.want)
			}
		})
	}
}
