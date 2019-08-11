package main

import "fmt"

func main() {
	s := make([]string, 4)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)

	fmt.Println("len", len(s))

	s = append(s, "d", "e")

	fmt.Println(append(s, "3"))

	fmt.Println(s)

	c := make([]string, len(s))

	copy(c, s)

	// slice opperator [low: high]

	l := c[2:5]
	fmt.Println("slice 2:5", l)

	// up to but excluding 5

	l = c[:5]
	fmt.Println("up to 5", l)

	l = c[2:]
	fmt.Println("from 2", l)

	d := []string{"pizza", "pie", "man"}

	fmt.Print(d)

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
}
