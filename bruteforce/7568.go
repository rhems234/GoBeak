package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	x := make([]int, n)
	y := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &x[i], &y[i])
	}

	for i := 0; i < n; i++ {
		rank[i]++
		for j := 0; j < n; j++ {
			if x[i] < x[j] && y[i] < y[j] {
				rank[i]++
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%d ", rank[i])
	}
	fmt.Fprintf(writer, "\n")
}
