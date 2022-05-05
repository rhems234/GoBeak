package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, c float64
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for true {
		fmt.Fscanln(reader, &a, &b, &c)

		if a == 0 && b == 0 && c == 0 {
			break
		}

		ausar := math.Pow(a, 2)
		auset := math.Pow(b, 2)
		heru := math.Pow(c, 2)

		result := "wrong"
		if ausar > auset && ausar > heru {
			if ausar-auset-heru == 0 {
				result = "right"
			}
		} else if auset > ausar && auset > heru {
			if auset-ausar-heru == 0 {
				result = "right"
			}
		} else {
			if heru-ausar-auset == 0 {
				result = "right"
			}
		}
		fmt.Fprintln(writer, result)
	}

}
