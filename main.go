package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	s := make([]string, n)
	for i := range s {
		s[i] = nextLine()
	}

	// slice copy
	b := append([]string{}, s...)
	b = bubbleSort(n, b)
	fmt.Println(strings.Join(b, " "))
	fmt.Println(isStable(n, s, b))

	// slice copy
	se := append([]string{}, s...)
	se = selectionSort(n, se)
	fmt.Println(strings.Join(se, " "))
	fmt.Println(isStable(n, s, se))
}

func isStable(n int, in, out []string) string {
	for i := 0; i <= n-1; i++ {
		for j := i + 1; j <= n-1; j++ {
			for a := 0; a <= n-1; a++ {
				for b := a + 1; b <= n-1; b++ {
					if in[i][1:] == in[j][1:] && in[i] == out[b] && in[j] == out[a] {
						return "Not stable"
					}
				}
			}
		}
	}
	return "Stable"
}

func bubbleSort(n int, s []string) []string {
	f := true
	for f {
		f = false
		for i := n - 1; 1 <= i; i-- {
			if s[i][1:] < s[i-1][1:] {
				v := s[i]
				s[i] = s[i-1]
				s[i-1] = v
				f = true
			}
		}
	}

	return s
}

func selectionSort(n int, s []string) []string {
	for i := range s {
		min := i
		for j := i; j <= n-1; j++ {
			if s[j][1:] < s[min][1:] {
				min = j
			}
		}
		if s[i][1:] != s[min][1:] {
			s[i], s[min] = swap(s[i], s[min])
		}
	}

	return s
}

func swap(a, b string) (string, string) {
	return b, a
}

func convertIntToString(s []int, d string) string {
	b := make([]string, len(s))
	for i, v := range s {
		b[i] = strconv.Itoa(v)
	}

	return strings.Join(b, d)
}

// =================================
// I/O util
// =================================

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// =================================
