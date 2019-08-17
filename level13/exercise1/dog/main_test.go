package dog

import (
	"fmt"
	"testing"

	"github.com/amassarwi/go-playground/level12/dog"
)

func TestYears(t *testing.T) {
	age := Years(10)
	if age != 70 {
		t.Error("Expected:", 70, "got", age)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func TestYearsTwo(t *testing.T) {
	age := YearsTwo(20)
	if age != 140 {
		t.Error("Expected:", 140, "got", age)
	}
}

func ExampleYears() {
	fmt.Println(dog.Years(10))
	// outputs:
	// 70
}
