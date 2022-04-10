package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var a int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &a)
	defer writer.Flush()

	for i := 0; i < a; i++ {
		var num int
		fmt.Fscanf(reader, "%d", &num)

		var score = make([]float64, num)
		var sum, avg float64
		for j := 0; j < num; j++ {
			fmt.Fscanf(reader, "%f", &score[j])
			sum = sum + score[j]
		}
		avg = sum / float64(num)

		var p float64
		for _, val := range score {
			if val > avg {
				p += (1 / float64(num))
			}
		}
		fmt.Fprintf(writer, "%.3f%s\n", math.Round(p*100000)/1000, "%")
	}

}
