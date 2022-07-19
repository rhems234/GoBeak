package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var input string
		fmt.Fscanln(reader, &input)
		fmt.Fprintln(writer, vps(input))
	}
}

func vps(input string) string {
	var open, close int

	for _, v := range input {
		if string(v) == "(" {
			open++
		} else if string(v) == ")" {
			close++
		}
		if open < close {
			return "NO"
		}
	}
	if open != close {
		return "NO"
	}
	return "YES"
}
