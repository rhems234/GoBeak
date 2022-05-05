package main

import "fmt"

func main() {
	food := make([]int, 3)
	for i := range food {
		fmt.Scanln(&food[i])
	}

	coke := make([]int, 2)
	for i := range coke {
		fmt.Scanln(&coke[i])
	}

	var minfood, mincoke = 2000, 2000

	for _, foods := range food {
		if foods < minfood {
			minfood = foods
		}
	}

	for _, cokes := range coke {
		if cokes < mincoke {
			mincoke = cokes
		}
	}

	fmt.Println(minfood + mincoke - 50)

}
