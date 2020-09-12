package main

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{s: "III"},
			want: 3,
		},
		{
			name: "case2",
			args: args{s: "IV"},
			want: 4,
		},
		{
			name: "case3",
			args: args{s: "IX"},
			want: 9,
		},
		{
			name: "case4",
			args: args{s: "LVIII"},
			want: 58,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
