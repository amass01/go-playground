package main

import "fmt"

func main() {

	m := []int{777, 209, 14, 13, 12, 11, 24, 13332, 122, 23232}

	for i, v := range m {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", m)

}
