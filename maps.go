package fn

import (
	"cmp"
	"slices"
)

// Entry represents a key-value pair in a map.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// GetOptional returns the value associated with the key in the map m. If the key is not present, it returns the default value.
func GetOptional[M ~map[K]V, K comparable, V any](m M, key K, defaultValue V) V {
	v, ok := m[key]
	if !ok {
		return defaultValue
	}
	return v
}

// Keys returns the keys of the map m. The keys will be in an indeterminate order.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// SortedKeys returns the keys of the map m in sorted order.
func SortedKeys[M ~map[K]V, K cmp.Ordered, V any](m M) []K {
	keys := Keys(m)
	slices.Sort(keys)
	return keys
}

// SortedKeysFunc returns the keys of the map m in sorted order using the given comparison function.
func SortedKeysFunc[M ~map[K]V, K comparable, V any](m M, cmp func(a, b K) int) []K {
	keys := Keys(m)
	slices.SortFunc(keys, cmp)
	return keys
}

// Values returns the values of the map m. The values will be in an indeterminate order.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Entries returns the (key, value) pairs of the map m. The pairs will be in an indeterminate order.
func Entries[M ~map[K]V, K comparable, V any](m M) []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(m))
	for k, v := range m {
		entries = append(entries, Entry[K, V]{k, v})
	}
	return entries
}

// Merge returns a new map that combines the contents of the original and the contents of the given maps.
func Merge[M ~map[K]V, K comparable, V any](original M, overrides ...M) M {
	result := make(M, len(original))
	for k, v := range original {
		result[k] = v
	}
	for _, m := range overrides {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// MergeInPlace adds the contents of the given maps to the original.
func MergeInPlace[M ~map[K]V, K comparable, V any](original M, overrides ...M) {
	for _, m := range overrides {
		for k, v := range m {
			original[k] = v
		}
	}
}
