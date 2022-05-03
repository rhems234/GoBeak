package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n, &m)

	number := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &number[i])
	}

	var sum int
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				var val = number[i] + number[j] + number[k]
				if val <= m && sum < val {
					sum = val
				} else {
					continue
				}
			}
		}
	}
	fmt.Println(sum)
}
