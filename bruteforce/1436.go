package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var number int
	var count int

	for true {
		number++
		var tmp = number
		var found = false
		for true {
			if tmp == 0 {
				break
			}
			if tmp%1000 == 666 {
				count++
				if count == n {
					found = true
				}
				break
			}
			tmp /= 10
		}
		if found {
			fmt.Println(number)
			break
		}
	}
}
