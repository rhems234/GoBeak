package main

import "fmt"

func main() {
	var n, m int
	var result int
	fmt.Scanln(&n, &m)

	result = (n * m) - 1

	fmt.Println(result)
}
