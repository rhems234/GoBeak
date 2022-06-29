package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var e = 1
	var s = 1
	var m = 1
	var E, S, M int
	var year = 1
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &E)
	fmt.Fscan(reader, &S)
	fmt.Fscan(reader, &M)

	for true {
		if e == E && s == S && m == M {
			fmt.Fprintln(writer, year)
			break
		}
		year++

		e += 1
		s += 1
		m += 1
		if e == 16 {
			e = 1
		}
		if s == 29 {
			s = 1
		}
		if m == 20 {
			m = 1
		}
	}
}
