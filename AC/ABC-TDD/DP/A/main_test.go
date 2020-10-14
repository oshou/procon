package a

import "testing"

func TestFrogA(t *testing.T) {
	type args struct {
		n int
		h []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{n: 4, h: []int{10, 30, 40, 20}},
			want: 30,
		},
		{
			name: "case2",
			args: args{n: 2, h: []int{10, 10}},
			want: 0,
		},
		{
			name: "case3",
			args: args{n: 6, h: []int{30, 10, 60, 10, 60, 50}},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FrogA(tt.args.n, tt.args.h); got != tt.want {
				t.Errorf("FrogA() = %v, want %v", got, tt.want)
			}
		})
	}
}
