package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var n = scanner.Text()

	var result = 0
	var tmp = 0
	var plus = true
	for _, s := range n {
		if i, err := strconv.Atoi(string(s)); err == nil {
			tmp = tmp*10 + i
			continue
		}

		if plus {
			result += tmp
		} else {
			result -= tmp
		}
		tmp = 0

		if plus == true && s == '-' {
			plus = false
		}
	}

	if plus {
		result += tmp
	} else {
		result -= tmp
	}
	fmt.Print(result)
}
