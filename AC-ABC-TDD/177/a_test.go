package a

import "testing"

func TestIsInTime(t *testing.T) {
	type args struct {
		d int
		t int
		s int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				d: 1000,
				t: 15,
				s: 80,
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				d: 2000,
				t: 20,
				s: 100,
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				d: 10000,
				t: 1,
				s: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInTime(tt.args.d, tt.args.t, tt.args.s); got != tt.want {
				t.Errorf("IsInTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
