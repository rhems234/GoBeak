package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	min := 100
	sum := 0
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 7; i++ {
		fmt.Fscanln(reader, &n)

		if n%2 == 1 {
			sum += n
			if n < min {
				min = n
			}
		}
	}
	if sum == 0 {
		fmt.Println("-1")
	} else {
		fmt.Println(sum, min)
	}
}
