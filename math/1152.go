package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')

	word := strings.Split(input, " ")
	var count int

	for i := range word {
		if word[i] != "\n" && word[i] != "" {
			count++
		}
	}
	fmt.Println(count)
}
