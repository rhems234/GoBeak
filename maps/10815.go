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

	var n int
	fmt.Fscanln(reader, &n)

	card := map[int]int{}
	for i := 0; i < n; i++ {
		var input int
		if i == n-1 {
			fmt.Fscanln(reader, &input)
		} else {
			fmt.Fscan(reader, &input)
		}
		card[input]++
	}

	var m int
	fmt.Fscanln(reader, &m)

	for i := 0; i < m; i++ {
		var num int
		fmt.Fscanf(reader, "%d", &num)
		fmt.Fprintf(writer, "%s ", hasCard(card, num))
	}
}

func hasCard(card map[int]int, num int) string {
	if card[num] != 0 {
		return "1"
	}
	return "0"
}
