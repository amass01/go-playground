package main

import "fmt"

func main() {
	const (
		today        = 2019
		nextYear     = today + iota
		inTwoYears   = today + iota
		inThreeYears = today + iota
		inFourYears  = today + iota
	)

	fmt.Println(today, nextYear, inTwoYears, inThreeYears, inFourYears)
}
