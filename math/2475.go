package main

import "fmt"

func main() {
	var n int
	var sum = 0

	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &n)
		sum += n * n
	}
	sum %= 10

	fmt.Printf("%d", sum)
}
