package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 19

	fmt.Println("map", m)
	v1 := m["k1"]

	fmt.Println(v1)

	fmt.Println(len(m))

	delete(m, "k2")
	fmt.Println("map", m)

	_, prs := m["k2"]

	fmt.Println(prs)

	mapping := map[string]int{"key1": 1, "key2": 2}

	fmt.Println(mapping)

}
