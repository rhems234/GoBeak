package main

import "fmt"

func bino(n, m int) (result int) {
	if n == m || m == 0 {
		return 1
	}
	return bino(n-1, m-1) + bino(n-1, m)
}

func main() {
	var n, m int
	fmt.Scanln(&n, &m)

	fmt.Println(bino(n, m))
}
