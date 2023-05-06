package unittest

import "testing"

func TestHomeworkBE4678487(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Length more than 26",
			args: args{
				input: "abcdefghijklmnopqrstuvwxyz1234567890",
			},
			want: -1,
		},
		{
			name: "Result on Success",
			args: args{
				input: "helloworld",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HomeworkBE4678487(tt.args.input); got != tt.want {
				t.Errorf("HomeworkBE4678487() = %v, want %v", got, tt.want)
			}
		})
	}
}
