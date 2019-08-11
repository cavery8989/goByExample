package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp", a)

	a[2] = 3
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// len gets me the length

	fmt.Println("length", len(a))

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println("dlc:", b)

	var c [2][3]int

	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			c[i][j] = i + j
		}
	}

	fmt.Println(c)
}
