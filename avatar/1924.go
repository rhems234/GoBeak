package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var month, date int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d", &month, &date)

	var days = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var day = [7]string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}

	for i := 0; i < month; i++ {
		date += days[i]
	}
	fmt.Printf("%s\n", day[date%7])
}
