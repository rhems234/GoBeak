package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n, &m)

	var mat = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%v\n", &mat[i])
	}
	var min = n * m
	for i := 0; i < n-7; i++ {
		for j := 0; j < m-7; j++ {
			var cnt1 = float64(0)
			var cnt2 = float64(0)
			for k := i; k < i+8; k++ {
				for l := j; l < j+8; l++ {
					if (k+l)%2 == 0 {
						if string(mat[k][l]) == "B" {
							cnt1++
						} else {
							cnt2++
						}
					} else {
						if string(mat[k][l]) == "B" {
							cnt2++
						} else {
							cnt1++
						}
					}
				}
			}
			if min > int(math.Min(cnt1, cnt2)) {
				min = int(math.Min(cnt1, cnt2))
			}
		}
	}
	fmt.Println(min)
}
