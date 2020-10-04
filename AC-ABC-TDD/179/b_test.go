package b

import "testing"

func Test_goToJail(t *testing.T) {
	type args struct {
		n int
		d [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{n: 5, d: [][]int{{1, 2}, {6, 6}, {4, 4}, {3, 3}, {3, 2}}},
			want: true,
		},
		{
			name: "case2",
			args: args{n: 5, d: [][]int{{1, 1}, {2, 2}, {3, 4}, {5, 5}, {6, 6}}},
			want: false,
		},
		{
			name: "case2",
			args: args{n: 6, d: [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goToJail(tt.args.n, tt.args.d); got != tt.want {
				t.Errorf("goToJail() = %v, want %v", got, tt.want)
			}
		})
	}
}
