package main

import (
	"fmt"
	"sort"
)

func main() {
	number := make([]int, 3)
	fmt.Scanf("%d %d %d", &number[0], &number[1], &number[2])
	sort.Slice(number, func(i, j int) bool {
		return number[i] < number[j]
	})

	fmt.Println(number[1])
}
