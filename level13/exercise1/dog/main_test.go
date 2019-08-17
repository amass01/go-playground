package dog

import "testing"

func TestYears(t *testing.T) {
	age := Years(10)
	if age != 70 {
		t.Error("Expected:", 70, "got", age)
	}
}

func TestYearsTwo(t *testing.T) {
	age := YearsTwo(20)
	if age != 140 {
		t.Error("Expected:", 140, "got", age)
	}
}
