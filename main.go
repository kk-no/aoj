package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {}

// =================================
// I/O util
// =================================

var sc = bufio.NewScanner(os.Stdin)

// Scan required
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
