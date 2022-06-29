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

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	var result, que []int
	for i := 1; i <= n; i++ {
		que = append(que, i)
	}

	var index int
	for len(que) > 0 {
		index = (index + k - 1) % len(que)
		result = append(result, que[index])
		que = append(que[:index], que[index+1:]...)
	}

	for i, v := range result {
		if i == 0 {
			fmt.Fprintf(writer, "<%d", v)
		} else {
			fmt.Fprintf(writer, ", %d", v)
		}
	}
	fmt.Fprintln(writer, ">")
}
