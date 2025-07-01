package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["Bob"] = "Robert"
	fmt.Println(m)

	n := map[int]int{0: 1, 1: 2, 2: 3}
	for k, v := range n {
		fmt.Printf("key: %d, value: %d\n", k, v)
	}

	s1 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(len(s1))
	for i, v := range s1 {
		fmt.Printf("key: %d, value: %d\n", i, v)
	}
}
