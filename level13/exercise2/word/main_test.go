package word

import (
	"fmt"
	"testing"
)

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("This is just a test")
	}
}

func BenchmarkUse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("This is just a test")
	}
}

func TestUseCount(t *testing.T) {
	s := UseCount("This is")
	if s["This"] != 1 {
		t.Error(`Expected res["This"] =`, 1, "got", s["This"])
	}
	if s["is"] != 1 {
		t.Error(`Expected res["is"] =`, 1, "got", s["is"])
	}
}

func TestCount(t *testing.T) {
	s := Count("This is")
	if s != 2 {
		t.Error(`Expected :`, 2, "got", s)
	}

	s = Count("")
	if s != 0 {
		t.Error(`Expected :`, 0, "got", s)
	}
}

func ExampleCount() {
	fmt.Println(Count("this is a test string"))
	// outpus:
	// 5
}
