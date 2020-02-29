package main

import "bufio"

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func main() {
	a := readString()
	m := readInt()
	h := make([]int, m)
	for i := 0; i < m; i++ {
		h = append(h, readInt())
	}
}
aabc
caab
aabc
bcaa
caab



xyzvw
vwxyz
zvwxy
