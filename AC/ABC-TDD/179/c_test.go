package solution

import "testing"

func Test_axb_c(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{n: 3},
			want: 3,
		},
		{
			name: "case2",
			args: args{n: 100},
			want: 473,
		},
		{
			name: "case3",
			args: args{n: 1000000},
			want: 13969985,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := axb_c(tt.args.n); got != tt.want {
				t.Errorf("axb_c() = %v, want %v", got, tt.want)
			}
		})
	}
}
