package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &input)

	sum := 0
	var key = []int{2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 10, 10, 10, 10}

	for i := 0; i < len(input); i++ {
		sum += key[int(input[i])-62]
	}
	fmt.Println(sum)
}
