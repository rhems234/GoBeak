package main

import "fmt"

func main() {
	var r, s int
	var r2 int
	fmt.Scanln(&r, &s)

	r2 = (s * 2) - r

	fmt.Println(r2)
}
