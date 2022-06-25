package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t, n int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)

		count := make([]int, n+1)

		if n+1 > 1 {
			count[1] = 1
		}
		if n+2 > 2 {
			count[2] = 2
		}
		if n+3 > 3 {
			count[3] = 4
		}
		for i := 4; i < n+1; i++ {
			count[i] = count[i-1] + count[i-2] + count[i-3]
		}
		fmt.Fprintln(writer, count[n])
	}

}
