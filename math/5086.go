package main

import "fmt"

func main() {
	var a, b int

	for i := 0; i < 10000; i++ {
		fmt.Scanln(&a, &b)

		if a == 0 && b == 0 {
			break
		}
		if a < b && b%a == 0 {
			fmt.Println("factor")
		} else if a > b && a%b == 0 {
			fmt.Println("multiple")
		} else {
			fmt.Println("neither")
		}
	}
}
