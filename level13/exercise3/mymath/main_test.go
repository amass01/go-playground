package mymath

import (
	"fmt"
	"testing"
)

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 5, 55, 123, 5555, 1222})
	}
}

func TestCenteredAvg(t *testing.T) {

	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{[]int{21, 5, 5, 5, 21}, 10.333333333333334},
		test{[]int{3, 8, 32, 20, 5}, 11},
	}

	for _, v := range tests {
		x := CenteredAvg(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3, 4}))
	// outputs:
	// 2.5
}
