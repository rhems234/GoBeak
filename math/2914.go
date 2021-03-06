package main

import (
	"bufio"
	"fmt"
	"os"
)

func LCM(a, b int) (lcm int) { // 최소공배수
	return a * b / GCD(a, b)
}

func GCD(a, b int) (gcd int) { // 최대공약수
	if a < b {
		b, a = a, b
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var a, b int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &a, &b)

	fmt.Println(GCD(a, b))
	fmt.Println(LCM(a, b))
}
