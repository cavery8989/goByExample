package main

import (
	"fmt"
	"math"
)

const s string = "potato"

func main() {

	fmt.Println(s)

	const num = 500000

	const d = 3e20 / num
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(num))

}
