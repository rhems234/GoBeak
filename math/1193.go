package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var level = 1
	var num = 1

	for true {
		if num > n {
			break
		}
		level++
		for i := level; i > 0; i-- {
			num++
			if num == n {
				if level%2 == 1 {
					fmt.Println(strconv.Itoa(i) + "/" + strconv.Itoa(level+1-i))
				} else {
					fmt.Println(strconv.Itoa(level+1-i) + "/" + strconv.Itoa(i))
				}
			}
		}
	}
}
