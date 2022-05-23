package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	var count = NumberCount(n)
	fmt.Println(count)

}

func NumberCount(number int) (count int) {
	if number < 100 {
		count = number
		return
	}
	for i := 100; i <= number; i++ {
		a := i % 10
		b := i / 10 % 10
		c := i / 100
		if c-b == b-a {
			count++
		}
	}
	count += 99
	return count
}
