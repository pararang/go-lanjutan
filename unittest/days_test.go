package unittest

import "testing"

func TestToDays(t *testing.T) {
	type args struct {
		amount int
		unit   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "should return -1 if amount is less than 1",
			args: args{amount: 0, unit: "w"},
			want: -1,
		},
		{
			name: "should return -1 if unit is empty",
			args: args{amount: 100, unit: ""},
			want: -1,
		},
		{
			name: "should return -1 if unit is not d or m or w",
			args: args{amount: 10, unit: "q"},
			want: -1,
		},
		{
			name: "should return amount if unit is d",
			args: args{amount: 10, unit:"d"},
			want: 10,
		},
		{
			name: "should return amount times 7 if unit is w",
			args: args{amount: 10, unit:"w"},
			want: 70,
		},
		{
			name: "should return amount times 30 if unit is m",
			args: args{amount: 10, unit:"m"},
			want: 300,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToDays(tt.args.amount, tt.args.unit); got != tt.want {
				t.Errorf("ToDays() = %v, want %v", got, tt.want)
			}
		})
	}
}


// package mytesting

// import "testing"

// func TestToDays(t *testing.T) {
// 	t.Run("should return -1 if amount is less than 1", func(t *testing.T) {
// 		amount := 0
// 		got := ToDays(amount, "w")
// 		want := -1

// 		if got != want {
// 			t.Errorf("ToDays() = %v, want %v", got, want)
// 		}
// 	})
	
// 	t.Run("should return -1 if amount is greater than 600", func(t *testing.T) {
// 		amount := 601
// 		got := ToDays(amount, "w")
// 		want := -1

// 		if got != want {
// 			t.Errorf("ToDays() = %v, want %v", got, want)
// 		}
// 	})
	
// 	t.Run("should return -1 if unit is empty", func(t *testing.T) {
// 		unit := ""
// 		got := ToDays(2, unit)
// 		want := -1

// 		if got != want {
// 			t.Errorf("ToDays() = %v, want %v", got, want)
// 		}
// 	})
// }