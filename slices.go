package fn

// At returns the element at the specified index in the slice.
// Negative index count back from the last item in the slice.
func At[S ~[]T, T any](slice S, index int) T {
	if index < 0 {
		index += len(slice)
	}
	return slice[index]
}

// AtOptional returns the element at the specified index in the slice or a default value if the index is out of bounds.
// Negative index count back from the last item in the slice.
func AtOptional[S ~[]T, T any](slice S, index int, defaultValue T) T {
	if index < 0 {
		index += len(slice)
	}
	if index < 0 || index >= len(slice) {
		return defaultValue
	}
	return slice[index]
}

// Map applies a function to each element in the slice and returns a new slice with the results.
func Map[S ~[]T, T any, U any](slice S, f func(value T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

// MapWithIndex applies a function to each element in the slice and returns a new slice with the results.
func MapWithIndex[S ~[]T, T any, U any](slice S, f func(value T, index int) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v, i)
	}
	return result
}

// MapInPlace applies a function to each element in the slice and modifies the slice in place.
func MapInPlace[S ~[]T, T any](slice S, f func(value T) T) {
	for i, v := range slice {
		slice[i] = f(v)
	}
}

// MapInPlaceWithIndex applies a function to each element in the slice and modifies the slice in place.
func MapInPlaceWithIndex[S ~[]T, T any](slice S, f func(value T, index int) T) {
	for i, v := range slice {
		slice[i] = f(v, i)
	}
}

// Reduce applies a function against an accumulator and each element in the slice (from left to right) to reduce it to a single value.
func Reduce[S ~[]T, T any, U any](s S, init U, f func(accumulator U, currentValue T) U) U {
	for _, v := range s {
		init = f(init, v)
	}
	return init
}

// ReduceWithIndex applies a function against an accumulator and each element in the slice (from left to right) to reduce it to a single value.
func ReduceWithIndex[S ~[]T, T any, U any](s S, init U, f func(accumulator U, currentValue T, currentIndex int) U) U {
	for i, v := range s {
		init = f(init, v, i)
	}
	return init
}

// ReduceRight applies a function against an accumulator and each element in the slice (from right to left) to reduce it to a single value.
func ReduceRight[S ~[]T, T any, U any](s S, init U, f func(accumulator U, currentValue T) U) U {
	for i := len(s) - 1; i >= 0; i-- {
		init = f(init, s[i])
	}
	return init
}

// ReduceRightWithIndex applies a function against an accumulator and each element in the slice (from right to left) to reduce it to a single value.
func ReduceRightWithIndex[S ~[]T, T any, U any](slice S, init U, f func(accumulator U, currentValue T, currentIndex int) U) U {
	for i := len(slice) - 1; i >= 0; i-- {
		init = f(init, slice[i], i)
	}
	return init

}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func Filter[S ~[]T, T any](slice S, f func(value T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Fill changes all elements in the slice to a static value.
func Fill[S ~[]T, T any](slice S, value T) {
	for i := range slice {
		slice[i] = value
	}
}

// Any test whether at least one element in the slice passes the test implemented by the provided function.
func Any[S ~[]T, T any](slice S, f func(value T) bool) bool {
	for _, v := range slice {
		if f(v) {
			return true
		}
	}
	return false
}

// Every test whether all elements in the slice pass the test implemented by the provided function.
func Every[S ~[]T, T any](slice S, f func(value T) bool) bool {
	for _, v := range slice {
		if !f(v) {
			return false
		}
	}
	return true
}
