package main

import "fmt"

func main() {
	favSport := "football"
	switch favSport {
	case "tennis":
		fmt.Println("not accurate!")
	case "football":
		fmt.Println("Korrekt - I;m in love with football")
	}

}
