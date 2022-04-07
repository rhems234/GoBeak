package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &a)

	for true {
		if a%5 != 0 {
			if a < 3 {
				b = -1
				break
			}
			a = a - 3
			b++
		} else {
			break
		}
	}
	if b != -1 {
		b += a / 5
	}
	fmt.Println(b)

}
