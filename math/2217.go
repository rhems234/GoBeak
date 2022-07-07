package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var rope = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &rope[i])
	}
	sort.Ints(rope)

	var max int
	for i := len(rope) - 1; i >= 0; i-- {
		weight := rope[i] * (len(rope) - i)
		if weight > max {
			max = weight
		}
	}
	fmt.Fprintln(writer, max)
}
