package abc163a

import (
	"math"
	"testing"
)

func Circumference(r float64) float64 {
	return 2 * r * math.Pi
}

func TestCircumference(t *testing.T) {
	cases := []struct {
		round float64
		want  float64
	}{
		{1, 6.28318530717958623200},
		{73, 458.67252742410977361942},
	}

	for _, tc := range cases {
		got := Circumference(tc.round)
		if got != tc.want {
			t.Errorf("failed. got=%v, want=%v", got, tc.want)
		}
	}

}
