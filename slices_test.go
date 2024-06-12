package fn

import (
	"slices"
	"testing"
)

func TestAt(t *testing.T) {
	slice := []int{1, 2, 3}
	if got, want := At(slice, 1), 2; got != want {
		t.Errorf("At(slice, 1) = %v, want %v", got, want)
	}
	if got, want := At(slice, -1), 3; got != want {
		t.Errorf("At(slice, -1) = %v, want %v", got, want)
	}
	if got, want := At(slice, -2), 2; got != want {
		t.Errorf("At(slice, -2) = %v, want %v", got, want)
	}
	{
		defer func() {
			if recover() == nil {
				t.Errorf("At(slice, 3) did not panic")
			}
		}()
		At(slice, 3)
	}
}

func TestAtOptional(t *testing.T) {
	slice := []int{1, 2, 3}
	if got, want := AtOptional(slice, 1, -1), 2; got != want {
		t.Errorf("At(slice, 1) = %v, want %v", got, want)
	}
	if got, want := AtOptional(slice, -1, -1), 3; got != want {
		t.Errorf("At(slice, -1) = %v, want %v", got, want)
	}
	if got, want := AtOptional(slice, -2, -1), 2; got != want {
		t.Errorf("At(slice, -2) = %v, want %v", got, want)
	}
	if got, want := AtOptional(slice, 3, -1), -1; got != want {
		t.Errorf("At(slice, 3) = %v, want %v", got, want)
	}
}

func TestMap(t *testing.T) {
	slice := []int{1, 2, 3}
	if got, want := Map(slice, func(v int) int { return v * 2 }), []int{2, 4, 6}; !slices.Equal(got, want) {
		t.Errorf("Map(slice, f) = %v, want %v", got, want)
	}
}

func TestMapWithIndex(t *testing.T) {
	slice := []int{1, 2, 3}
	if got, want := MapWithIndex(slice, func(v int, i int) int { return v + i }), []int{1, 3, 5}; !slices.Equal(got, want) {
		t.Errorf("MapWithIndex(slice, f) = %v, want %v", got, want)
	}
}

func TestMapInPlace(t *testing.T) {
	slice := []int{1, 2, 3}
	MapInPlace(slice, func(v int) int { return v * 2 })
	if got, want := slice, []int{2, 4, 6}; !slices.Equal(got, want) {
		t.Errorf("MapInPlace(slice, f) = %v, want %v", got, want)
	}
}

func TestMapInPlaceWithIndex(t *testing.T) {
	slice := []int{1, 2, 3}
	MapInPlaceWithIndex(slice, func(v int, i int) int { return v + i })
	if got, want := slice, []int{1, 3, 5}; !slices.Equal(got, want) {
		t.Errorf("MapInPlaceWithIndex(slice, f) = %v, want %v", got, want)
	}
}

func TestReduce(t *testing.T) {
	slice := []int{1, 2, 3}
	if got, want := Reduce(slice, 0, func(a, b int) int { return a + b }), 6; got != want {
		t.Errorf("Reduce(slice, f) = %v, want %v", got, want)
	}
}

func TestReduceWithIndex(t *testing.T) {
	slice := []int{1, 2, 3}
	if got, want := ReduceWithIndex(slice, 0, func(a, b int, i int) int { return a + b + i }), 9; got != want {
		t.Errorf("ReduceWithIndex(slice, f) = %v, want %v", got, want)
	}
}

func TestReduceRight(t *testing.T) {
	slice := []int{1, 2, 3}
	if got, want := ReduceRight(slice, 0, func(a, b int) int { return a + b }), 6; got != want {
		t.Errorf("ReduceRight(slice, f) = %v, want %v", got, want)
	}
}

func TestReduceRightWithIndex(t *testing.T) {
	slice := []int{1, 2, 3}
	if got, want := ReduceRightWithIndex(slice, 0, func(a, b int, i int) int { return a + b + i }), 9; got != want {
		t.Errorf("ReduceRightWithIndex(slice, f) = %v, want %v", got, want)
	}
}

func TestFilter(t *testing.T) {
	slice := []int{1, 2, 3}
	if got, want := Filter(slice, func(v int) bool { return v%2 == 0 }), []int{2}; !slices.Equal(got, want) {
		t.Errorf("Filter(slice, f) = %v, want %v", got, want)
	}
}

func TestFill(t *testing.T) {
	slice := make([]int, 3)
	Fill(slice, 1)
	if got, want := slice, []int{1, 1, 1}; !slices.Equal(got, want) {
		t.Errorf("Fill(slice, 1) = %v, want %v", got, want)
	}
}

func TestAny(t *testing.T) {
	slice := []int{1, 2, 3}
	if got, want := Any(slice, func(v int) bool { return v%2 == 0 }), true; got != want {
		t.Errorf("Any(slice, f) = %v, want %v", got, want)
	}

	if got, want := Any(slice, func(v int) bool { return v < 0 }), false; got != want {
		t.Errorf("Any(slice, f) = %v, want %v", got, want)
	}
}

func TestEvery(t *testing.T) {
	slice := []int{1, 2, 3}
	if got, want := Every(slice, func(v int) bool { return v%2 == 0 }), false; got != want {
		t.Errorf("Every(slice, f) = %v, want %v", got, want)
	}

	if got, want := Every(slice, func(v int) bool { return v > 0 }), true; got != want {
		t.Errorf("Every(slice, f) = %v, want %v", got, want)
	}
}
