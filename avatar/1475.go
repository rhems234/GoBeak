package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n)

	if n == 0 {
		fmt.Fprintln(writer, 1)
		return
	}

	number := map[int]int{}
	for n > 0 {
		number[n%10]++
		n /= 10
	}

	six, nine := number[6], number[9]
	if (six+nine)%2 == 0 {
		number[6] = (six + nine) / 2
		number[9] = number[6]
	} else {
		number[6] = (six + nine) / 2
		number[9] = number[6] + 1
	}

	var set int

	for _, v := range number {
		if set < v {
			set = v
		}
	}
	fmt.Fprintln(writer, set)
}
