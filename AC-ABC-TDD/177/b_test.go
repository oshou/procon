package b

import "testing"

func Test_substring(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				s: "cabacc",
				t: "abc",
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				s: "codeforces",
				t: "atcoder",
			},
			want: 6,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := substring(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("substring() = %v, want %v", got, tt.want)
			}
		})
	}
}
