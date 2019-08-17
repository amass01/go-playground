package main

import "fmt"

type customErr struct {
	message string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Something bad happened: %v", ce.message)
}

func foo(err error) {
	fmt.Println(err)
}

func main() {
	ce := customErr{
		message: "test me error",
	}

	foo(ce)
}
