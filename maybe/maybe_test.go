package maybe

import (
	"strings"
	"testing"

	"github.com/erikjuhani/go-fp/pipe"
)

func head[T any](x []T) Maybe[T] {
	if len(x) > 0 {
		return Just(x[0])
	}
	return Nothing[T]()
}

func inverse(x int) Maybe[float32] {
	if x > 0 {
		return Just(1 / float32(x))
	}

	return Nothing[float32]()
}

func TestMap(t *testing.T) {
	tests := []struct {
		expected string
		data     []string
	}{
		{},
		{"HELLO", []string{"hello", "world"}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := pipe.Pipe2(
				head[string],
				Map(strings.ToUpper),
			)(tt.data)

			res := Match(
				func() string { return "" },
				func(v string) string { return v },
			)(result)

			if res != tt.expected {
				t.Errorf("expected %s, but got %s", tt.expected, res)
			}

		})
	}
}

func TestFmap(t *testing.T) {
	tests := []struct {
		expected float32
		data     []int
	}{
		{},
		{0, []int{0, 10}},
		{0.2, []int{5, 10}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := pipe.Pipe2(
				head[int],
				Fmap(inverse),
			)(tt.data)

			res := Match(
				func() float32 { return 0 },
				func(v float32) float32 { return v },
			)(result)

			if res != tt.expected {
				t.Errorf("expected %f, but got %f", tt.expected, res)
			}
		})
	}
}

func TestFrom(t *testing.T) {
	var (
		x *int
		y = new(int)
	)
	*y = 0
	tests := []struct {
		expected int
		arg0     any
		arg1     bool
	}{
		{},
		{0, 1, false},
		{0, x, true},
		{0, y, true},
		{1, 1, true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := From(tt.arg0, tt.arg1)

			res := Match(
				func() any { return 0 },
				func(v any) any { return v },
			)(result)

			if res != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, res)
			}
		})
	}
}
