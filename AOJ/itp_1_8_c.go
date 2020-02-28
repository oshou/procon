package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanLines)
	sc.Buffer([]byte{}, math.MaxInt64)
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	m := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0,
		"m": 0,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 0,
		"t": 0,
		"u": 0,
		"v": 0,
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}
	s := readString()
	for i := 0; i < len(s); i++ {
		c := strings.ToLower(string(s[i]))
		if c != "." && c != " " {
			m[c]++
		}
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%v:%v\n", key, m[key])
	}
}
