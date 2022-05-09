package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var input string
	var sum int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &input)

		var letter = make(map[uint8]bool)
		var prev uint8
		var result = true
		for j := 0; j < len(input); j++ {
			_, exist := letter[input[j]]
			if !exist {
				letter[input[j]] = true
			} else if prev != input[j] {
				result = false
				break
			}
			prev = input[j]
		}
		if result {
			sum++
		}
	}
	fmt.Println(sum)
}
