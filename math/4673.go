package main

import "fmt"

func selfNumber(n int) (snumber map[int]bool) {
	snumber = make(map[int]bool, n+1)
	for i := 0; i < n+1; i++ {
		snumber[i] = false
	}

	for i := 0; i < n+1; i++ {
		var sum = i
		var number = i
		for j := number; j != 0; j /= 10 {
			sum += j % 10
		}
		if sum <= n {
			snumber[sum] = true
		}
	}

	return snumber
}

func main() {
	snumber := selfNumber(10000)
	for i := 0; i < len(snumber); i++ {
		if snumber[i] == false {
			fmt.Println(i)
		}
	}
}
