package b

import "testing"

func Test_numPoints(t *testing.T) {
	type args struct {
		n int
		d int
		x []int
		y []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n: 4,
				d: 5,
				x: []int{0, -2, 3, 4},
				y: []int{5, 4, 4, -4},
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				n: 12,
				d: 3,
				x: []int{1, 1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3},
				y: []int{1, 1, 1, 1, 2, 3, 1, 2, 3, 1, 2, 3},
			},
			want: 7,
		},
		{
			name: "case3",
			args: args{
				n: 20,
				d: 100000,
				x: []int{14309, -56855, 151364, 103789, 147404, -37006, 188810, 13419, -88280, -196399, -176527, 46659, -153551, 98784, 94111, -30401, -55056, 5901, 138819, -69848},
				y: []int{-32939, 100340, 25430, -113141, -136977, -30929, -49557, 70401, 165170, 137941, -61904, 115261, 114185, -6820, -86268, 61477, 7872, -163796, -185986, -96669},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numPoints(tt.args.n, tt.args.d, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("numPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
