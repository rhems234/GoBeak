package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	number := make([]bool, m+1)

	for i := 2; i <= m; i++ {
		if number[i] {
			continue
		}
		if i >= n {
			fmt.Println(i)
		}
		for j := i * 2; j <= m; j += i {
			number[j] = true
		}
	}
}
