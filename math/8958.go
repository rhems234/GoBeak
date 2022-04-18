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

	var result = make([]string, n)
	var scores = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &result[i])

		var sum, score int
		for j := 0; j < len(result[i]); j++ {
			if fmt.Sprintf("%c", result[i][j]) == "o" {
				sum++
				score += sum
			} else {
				sum = 0
			}
		}
		scores[i] = score
	}
	for _, val := range scores {
		fmt.Println(val)
	}
}
