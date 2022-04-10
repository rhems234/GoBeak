package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var count int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d", &n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &i)
		if i%2 == 1 {
			count++
		} else if i == 1 {
			count--
		}
	}

	fmt.Println(count)
}
