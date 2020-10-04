package a

import (
	"fmt"
	"testing"
)

func Test_nums(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{s: "apple"},
			want: "apples",
		},
		{
			name: "case2",
			args: args{s: "bus"},
			want: "buses",
		},
		{
			name: "case3",
			args: args{s: "box"},
			want: "boxs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nums(tt.args.s); got != tt.want {
				fmt.Println(got)
				t.Errorf("nums() = %v, want %v", got, tt.want)
			}
		})
	}
}
