package main

import "fmt"

func main() {
	var t, n int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var sum = 0
		var min = 100
		for j := 0; j < 7; j++ {
			fmt.Scanf("%d", &n)
			if n%2 == 0 {
				sum += n

				if min > n {
					min = n
				}
			}
		}
		fmt.Printf("%d %d\n", sum, min)
	}
}
