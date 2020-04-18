package main

// Insertion Sort

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
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	for i, v := range s {
		for j := i - 1; 0 <= j && v < s[j]; j-- {
			s[j+1] = s[j]
			s[j] = v
		}
		fmt.Println(convertIntToString(s, " "))
	}
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
