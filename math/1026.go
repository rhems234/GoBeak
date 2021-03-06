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

	a := make([]int, n)
	b := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &b[i])
	}

	sort.Ints(a)
	sort.Slice(b, func(i, j int) bool {
		return b[j] < b[i]
	})

	var result int
	for i := 0; i < n; i++ {
		result += a[i] * b[i]
	}
	fmt.Fprintln(writer, result)
}
