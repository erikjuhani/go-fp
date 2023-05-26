package state

import (
	"testing"

	"github.com/erikjuhani/go-fp/pipe"
)

func TestPut(t *testing.T) {
	var (
		incrementCounterState = Fmap(func(s int) State[Void, int] { return Put(s + 1) })
	)

	tests := []struct {
		expected     int
		initialState int
	}{
		{0, -1},
		{1, 0},
		{2, 1},
	}

	counter := pipe.Pipe3(
		Run[int],
		incrementCounterState,
		Exec[Void](0),
	)

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := counter(tt.initialState)

			if result != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, result)
			}
		})
	}
}

func TestModify(t *testing.T) {
	var (
		getCurrentCounterState = Fmap(Get[int])
		incrementCounterState  = Fmap(func(int) State[Void, int] { return Modify(func(s int) int { return s + 1 }) })
	)

	tests := []struct {
		expected     int
		initialState int
	}{
		{0, -1},
		{1, 0},
		{2, 1},
	}

	counter := pipe.Pipe4(
		Run[int],
		getCurrentCounterState,
		incrementCounterState,
		Exec[Void](0),
	)

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := counter(tt.initialState)

			if result != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, result)
			}
		})
	}
}

func TestGetS(t *testing.T) {
	tests := []struct {
		expected     int
		initialState int
	}{
		{0, -1},
		{1, 0},
		{2, 1},
	}

	increment := pipe.Pipe3(
		Run[int],
		Fmap(func(int) State[int, int] { return GetS(func(s int) int { return s + 1 }) }),
		Eval[int](0),
	)

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := increment(tt.initialState)

			if result != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, result)
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		expected     int
		initialState int
	}{
		{9, -1},
		{10, 0},
		{11, 1},
	}

	add := func(a int) func(b int) int {
		return func(b int) int {
			return a + b
		}
	}

	addTen := pipe.Pipe3(
		Run[int],
		Map[int](add(10)),
		Eval[int](0),
	)

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := addTen(tt.initialState)

			if result != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, result)
			}
		})
	}
}
