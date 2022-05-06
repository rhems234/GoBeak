package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var i = 2
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &n)

	for true {
		if n == 1 {
			break
		}

		if n%i == 0 {
			n /= i
			fmt.Println(i)
		} else {
			i++
		}
	}
}
