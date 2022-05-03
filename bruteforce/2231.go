package main

import (
	"bufio"
	"fmt"
	"os"
)

func resultsum(n int) (result int) {
	result = n
	for n != 0 {
		result += n % 10
		n /= 10
	}
	return
}

func main() {
	var num int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &num)

	var result int
	for i := 0; i < num; i++ {
		var sum = resultsum(i)
		if num == sum {
			result = i
			break
		}
	}
	fmt.Println(result)
}
