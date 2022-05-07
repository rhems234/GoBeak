package main

import "fmt"

func main() {
	var text string

	fmt.Scanf("%s", &text)

	for i := 0; i < len(text); i++ {
		fmt.Printf("%c", text[i])

		if i%10 == 9 {
			fmt.Print("\n")
		}
	}
}
