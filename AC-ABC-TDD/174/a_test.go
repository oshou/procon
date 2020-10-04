package a

import "testing"

func TestNeedsAirConditioner(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{x: 25},
			want: false,
		},
		{
			name: "case2",
			args: args{x: 30},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NeedsAirConditioner(tt.args.x); got != tt.want {
				t.Errorf("NeedsAirConditioner() = %v, want %v", got, tt.want)
			}
		})
	}
}
