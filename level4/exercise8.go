package main

import "fmt"

func main() {
	m := map[string][]string{
		"James":           []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, values := range m {
		fmt.Println(k)
		for i, v := range values {
			fmt.Println(i, v)
		}
	}
}
