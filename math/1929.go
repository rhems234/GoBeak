package main

import (
	"bufio"
	"fmt"
	"os"
)

func PrimeNumber(n, m int) (kth int) {
	Prime := initPrime(n)
	for i := 2; i < n+1; i++ {
		if Prime[i] {
			for j := i; j < n+1; j += i {
				if Prime[j] {
					Prime[j] = false
					m--
					if m == 0 {
						kth = j
						return
					}
				}
			}
		}
	}
	return
}

func initPrime(n int) []bool {
	var PrimeNumber = make([]bool, n+1)
	for i := range PrimeNumber {
		PrimeNumber[i] = true
	}
	return PrimeNumber
}
func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n, &m)
	fmt.Fprintln(writer, PrimeNumber(n, m))

}
