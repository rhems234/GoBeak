package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	number := make([]int, 0, n)

	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scanln(&tmp)
		number = append(number, tmp)
	}

	sort.Ints(number)

	for _, i := range number {
		fmt.Println(i)
	}
}
