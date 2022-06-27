package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n = [8]int{0}
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 8; i++ {
		fmt.Fscan(reader, &n[i])
	}

	var result = ""

	for i := 0; i < len(n)-1; i++ {
		if n[i] < n[i+1] && (result == "" || result == "ascending") {
			result = "ascending"
		} else if n[i] > n[i+1] && (result == "" || result == "descending") {
			result = "descending"
		} else {
			result = "mixed"
			break
		}
	}
	fmt.Println(result)
}
