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

	word := make(map[uint8]int)
	for i := 0; i < 26; i++ {
		word[uint8(i)+65] = 0
	}

	for i := 0; i < len(input); i++ {
		ascii := input[i]
		if ascii > 90 {
			ascii -= 32
		}
		word[ascii]++
	}

	var maxval = -1
	var maxkey string
	for key, val := range word {
		if val > maxval {
			maxval = val
			maxkey = string(key)
		} else if val == maxval {
			maxkey = "?"
		}
	}
	fmt.Println(maxkey)
}
