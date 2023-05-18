package pipe

import (
	"testing"
)

func double(x int) int {
	return x * 2
}

func amountAsDozenString(x int) string {
	switch x {
	case 12:
		return "dozen"
	case 6:
		return "half a dozen"
	default:
		return "less or more than a dozen"
	}
}

func TestPipe(t *testing.T) {
	expected := 4096
	result := Pipe12(
		double,
		double,
		double,
		double,
		double,
		double,
		double,
		double,
		double,
		double,
		double,
		double,
	)(1)

	if result != expected {
		t.Errorf("expected %d, but got %d", expected, result)
	}
}

func TestPipeChangeType(t *testing.T) {
	expected := "half a dozen"
	result := Pipe2(
		double,
		amountAsDozenString,
	)(3)

	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}
