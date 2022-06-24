package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n)
	number := make([]int, 0, n)

	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scanln(&tmp)
		number = append(number, tmp)
	}

	sort.Ints(number)

	for _, i := range number {
		fmt.Fprintln(writer, i)
	}
}
