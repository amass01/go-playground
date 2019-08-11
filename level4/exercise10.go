package main

import "fmt"

func main() {
	m := map[string][]string{
		"James":           []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["amir"] = []string{`programming`, `taveling`, `blockchain`}

	for k, values := range m {
		fmt.Println(k, values)
	}

	delete(m, "James")

	fmt.Println("") // just an empty line
	fmt.Println("After deleting 'James':")
	for k, values := range m {
		fmt.Println(k, values)
	}
}
