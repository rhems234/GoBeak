package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var t, n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	primeNumber := getPrimeNumber(10000)
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)
		goldbach(n, primeNumber, writer)
	}
}

func getPrimeNumber(num int) (primeNumber map[int]bool) {
	primeNumber = make(map[int]bool, num)
	for i := 0; i < num; i++ {
		primeNumber[i+1] = true
	}
	for i := 0; i < num; i++ {
		if primeNumber[i+1] == false {
			continue
		}
		var count int
		for j := 0; j < int(math.Sqrt(float64(i+1))); j++ {
			if (i+1)%(j+1) == 0 {
				count++
				if count > 1 {
					for k := 1; k*(i+1) <= num; k++ {
						primeNumber[k*(i+1)] = false
					}
					break
				}
			}
		}
	}
	return primeNumber
}

func goldbach(num int, primeNumber map[int]bool, writer *bufio.Writer) {
	for i := num / 2; i > 0; i-- {
		if primeNumber[i] && primeNumber[num-i] {
			fmt.Fprintf(writer, "%d %d\n", i, num-i)
			break
		}
	}
}
