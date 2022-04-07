package main

import "fmt"

func main() {
	var a, b, v int
	fmt.Scanln(&a, &b, &v)

	var sum = 1

	if (v-a)%(a-b) == 0 {
		sum += (v - a) / (a - b)
	} else {
		sum += (v-a)/(a-b) + 1
	}

	fmt.Println(sum)
}
