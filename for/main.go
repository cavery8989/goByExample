package main

import "fmt"

func main() {
	fmt.Println("yo")

	i := 1

	for i < 3 {
		fmt.Println(i)

		i = i + 1
	}

	for j := 0; j <= 10; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("hat")
		break
	}

	for k := 0; k <= 10; k++ {
		if k%2 != 0 {
			continue
		}
		fmt.Println(k)
	}
}
