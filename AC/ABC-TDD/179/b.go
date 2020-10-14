package b

func goToJail(n int, d [][]int) bool {
	var cnt, max int
	for _, di := range d {
		if di[0] == di[1] {
			cnt++
			if max < cnt {
				max = cnt
			}
			continue
		}
		cnt = 0
	}

	if max >= 3 {
		return true
	} else {
		return false
	}
}
