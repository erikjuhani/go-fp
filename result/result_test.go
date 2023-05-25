package result

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/erikjuhani/go-fp/pipe"
)

func head[T any](slice []T) Result[T] {
	if len(slice) > 0 {
		return Ok(slice[0])
	}
	return Err[T](errors.New("cannot get head from an empty array"))
}

func inverse(x int) Result[float32] {
	if x == 0 {
		return Err[float32](errors.New("division by zero"))
	}

	return Ok(1 / float32(x))
}

func TestMap(t *testing.T) {
	tests := []struct {
		expected string
		data     []string
	}{
		{"cannot get head from an empty array", []string{}},
		{"HELLO", []string{"hello", "world"}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := pipe.Pipe2(
				head[string],
				Map(strings.ToUpper),
			)(tt.data)

			res := Match(
				func(err error) string { return err.Error() },
				func(val string) string { return val },
			)(result)

			if res != tt.expected {
				t.Errorf("expected %s, but got %s", tt.expected, res)
			}

		})
	}
}

func TestFmap(t *testing.T) {
	tests := []struct {
		expected string
		data     []int
	}{
		{"cannot get head from an empty array", []int{}},
		{"division by zero", []int{0}},
		{"0.2", []int{5, 10}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := pipe.Pipe2(
				head[int],
				Fmap(inverse),
			)(tt.data)

			res := Match(
				func(err error) string { return err.Error() },
				func(val float32) string { return fmt.Sprintf("%.1f", val) },
			)(result)

			if res != tt.expected {
				t.Errorf("expected %s, but got %s", tt.expected, res)
			}
		})
	}
}

func TestFrom(t *testing.T) {
	tests := []struct {
		expected string
		arg0     string
		arg1     error
	}{
		{},
		{"Success", "Success", nil},
		{"Failure", "", errors.New("Failure")},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := From(tt.arg0, tt.arg1)

			res := Match(
				func(err error) string { return err.Error() },
				func(val string) string { return val },
			)(result)

			if res != tt.expected {
				t.Errorf("expected %s, but got %s", tt.expected, res)
			}
		})
	}
}

func TestUnwrap(t *testing.T) {
	tests := []struct {
		expected string
		data     []string
	}{
		{},
		{"", []string{}},
		{"Success", []string{"Success"}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := pipe.Pipe2(
				head[string],
				Unwrap[string],
			)(tt.data)

			if result != tt.expected {
				t.Errorf("expected %s, but got %s", tt.expected, result)
			}
		})
	}
}

func TestUnsafeUnwrap(t *testing.T) {
	tests := []struct {
		expected string
		data     []string
	}{
		{"cannot get head from an empty array", []string{}},
		{"Success", []string{"Success"}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			defer func() {
				r := recover()

				if r != nil && r.(error).Error() != errors.New(tt.expected).Error() {
					t.Errorf("expected %s, but got %s", tt.expected, r)
				}
			}()

			result := pipe.Pipe2(
				head[string],
				Unsafe_Unwrap[string],
			)(tt.data)

			if result != tt.expected {
				t.Errorf("expected %s, but got %s", tt.expected, result)
			}
		})
	}
}
