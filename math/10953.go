package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	var a, b int
	result := 0
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		fmt.Fscanf(reader, "%d,%d\n", &a, &b)

		result = a + b

		fmt.Println(result)
	}
}
