package main

import (
	"bufio"
	"fmt"
	"os"
)

func getcount(k, n int) (count int) {
	if k == 1 {
		for i := 0; i < n; i++ {
			count += (i + 1)
		}
		return count
	}
	for i := 0; i < n; i++ {
		count += getcount(k-1, i+1)
	}
	return count
}

func main() {
	var t, k, n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &t)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &k)
		fmt.Fscanln(reader, &n)
		fmt.Fprintln(writer, getcount(k, n))
	}
}
