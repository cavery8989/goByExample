package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	fmt.Println("sum: ", sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}

	// loop over range
	for k, v := range kvs {
		fmt.Printf("%s --> %s \n", k, v)
	}

	// can just loop over range.

	for k := range kvs {
		fmt.Println("key: ", k)
	}

	// can itereate over strings

	for i, r := range "go" {
		fmt.Println(i, r)
	}
}
