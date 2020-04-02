package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
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

type Pair struct {
	Key   string
	Value int
}

type Pairs []Pair

var _ sort.Interface = Pairs{}

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortMap(p map[string]int) Pairs {
	pairs := make(Pairs, len(p))
	i := 0
	for k, v := range p {
		pairs[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pairs))
	return pairs
}

func main() {
	n := readInt()
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	m := readInt()
	p := make(map[string]int)
	for i := 0; i < m; i++ {
		p[readString()] += readInt()
	}
	sorted := sortMap(p)
	for _, v := range sorted {
		fmt.Println(v.Key)
	}
}
