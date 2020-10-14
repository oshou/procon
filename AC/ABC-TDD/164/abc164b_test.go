package abc164b

import "testing"

func IsWinner(thp, tatk, ahp, aatk int) bool {
	tTurn := thp / aatk
	if thp%aatk != 0 {
		tTurn++
	}
	aTurn := ahp / tatk
	if ahp%tatk != 0 {
		aTurn++
	}
	if tTurn >= aTurn {
		return true
	}
	return false
}

func TestIsWinner(t *testing.T) {
	cases := []struct {
		ahp  int
		aatk int
		bhp  int
		batk int
		want bool
	}{
		{10, 9, 10, 10, false},
		{46, 4, 40, 5, true},
	}
	for _, tc := range cases {
		got := IsWinner(tc.ahp, tc.aatk, tc.bhp, tc.batk)
		if got != tc.want {
			t.Errorf("failed. got=%v, want=%v", got, tc.want)
		}
	}
}
