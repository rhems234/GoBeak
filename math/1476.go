package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var e, s, m int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &e, &s, &m)

	e %= 15
	s %= 28
	m %= 19

	var result int
	for i := 1; i > 0; i++ {
		if i%15 == e && i%28 == s && i%19 == m {
			result += i
		}
	}
	fmt.Print(result)
}
