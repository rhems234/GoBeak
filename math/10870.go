package main

import (
	"bufio"
	"fmt"
	"os"
)

func fibonacci(n int) (result int) {
	if n == 0 || n == 1 {
		result = n
		return result
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var a int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &a)

	var result = fibonacci(a)
	fmt.Println(result)

}
