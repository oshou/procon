package main

import "testing"

func Test_calc(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{a: 2},
			want: 14,
		},
		{
			name: "case2",
			args: args{a: 10},
			want: 1110,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calc(tt.args.a); got != tt.want {
				t.Errorf("calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
