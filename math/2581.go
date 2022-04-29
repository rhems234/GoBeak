package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &m)
	fmt.Fscanln(reader, &n)

	var sum = 0
	var min = n + 1
	for i := m; i < n+1; i++ {
		var count int

		for j := 0; j < i; j++ {
			if i%(j+1) == 0 {
				count++
			}
		}
		if count == 2 {
			sum += i
			if min > i {
				min = i
			}
		}
	}
	if sum == 0 {
		fmt.Println(-1)
		return
	}
	fmt.Printf("%d\n%d\n", sum, min)
}
