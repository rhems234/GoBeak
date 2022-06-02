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
	fmt.Fscanln(reader, &input)

	input = strings.Replace(input, "c=", "!", -1)
	input = strings.Replace(input, "c-", "@", -1)
	input = strings.Replace(input, "dz=", "#", -1)
	input = strings.Replace(input, "d-", "$", -1)
	input = strings.Replace(input, "lj", "%", -1)
	input = strings.Replace(input, "nj", "^", -1)
	input = strings.Replace(input, "s=", "&", -1)
	input = strings.Replace(input, "z=", "*", -1)

	var word []string
	for i := 0; i < len(input); i++ {
		word = append(word, string(input[i]))
	}
	fmt.Println(len(word))

}
