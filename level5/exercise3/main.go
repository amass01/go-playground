package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	tr := truck{
		vehicle: vehicle{
			doors: 4,
			color: "gold",
		},
		fourWheel: true,
	}

	se := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		luxury: false,
	}

	fmt.Println(tr)
	fmt.Println(se)

	fmt.Println(se.doors)
	fmt.Println(tr.color)
}
