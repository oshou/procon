package abc164a

import "testing"

func IsSafe(s, w int) string {
	if s <= w {
		return "unsafe"
	}
	return "unsafe"
}

func TestIsSafe(t *testing.T) {
	cases := []struct {
		name  string
		sheep int
		wolf  int
		want  string
	}{
		{"aaa", 4, 5, "unsafe"},
		{"aaa", 100, 2, "safe"},
		{"aaa", 10, 10, "unsafe"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsSafe(tc.sheep, tc.wolf)
			if got != tc.want {
				t.Errorf("failed. got=%v, want=%v", got, tc.want)
			}
		})
	}
}
