package fn

import (
	"cmp"
	"maps"
	"slices"
	"testing"
)

func TestGetOptional(t *testing.T) {
	m := map[int]string{1: "one", 2: "two"}
	if got, want := GetOptional(m, 1, "default"), "one"; got != want {
		t.Errorf("GetOptional(m, 1) = %v, want %v", got, want)
	}
	if got, want := GetOptional(m, 3, "default"), "default"; got != want {
		t.Errorf("GetOptional(m, 3) = %v, want %v", got, want)
	}
}

func TestKeys(t *testing.T) {
	m := map[int]string{1: "one", 2: "two"}
	got, want := Keys(m), []int{1, 2}
	slices.Sort(got)
	if !slices.Equal(got, want) {
		t.Errorf("Keys(m) = %v, want %v", got, want)
	}
}

func TestValues(t *testing.T) {
	m := map[int]string{1: "one", 2: "two"}
	got, want := Values(m), []string{"one", "two"}
	slices.Sort(got)
	if !slices.Equal(got, want) {
		t.Errorf("Values(m) = %v, want %v", got, want)
	}
}

func TestSortedKeys(t *testing.T) {
	m := map[int]string{2: "two", 1: "one"}
	if got, want := SortedKeys(m), []int{1, 2}; !slices.Equal(got, want) {
		t.Errorf("SortedKeys(m) = %v, want %v", got, want)
	}
}

func TestSortedKeysFunc(t *testing.T) {
	m := map[int]string{2: "two", 1: "one"}
	if got, want := SortedKeysFunc(m, cmp.Compare[int]), []int{1, 2}; !slices.Equal(got, want) {
		t.Errorf("SortedKeysFunc(m) = %v, want %v", got, want)
	}
}

func TestEntries(t *testing.T) {
	m := map[int]string{1: "one", 2: "two", 3: "three"}
	got, want := Entries(m), []Entry[int, string]{{1, "one"}, {2, "two"}, {3, "three"}}
	slices.SortFunc(got, func(a, b Entry[int, string]) int {
		return a.Key - b.Key
	})
	if !slices.Equal(got, want) {
		t.Errorf("Entries(m) = %v, want %v", got, want)
	}
}

func TestMerge(t *testing.T) {
	m1 := map[int]string{1: "one", 2: "two"}
	m2 := map[int]string{2: "deux", 3: "trois"}
	got, want := Merge(m1, m2), map[int]string{1: "one", 2: "deux", 3: "trois"}
	if !maps.Equal(got, want) {
		t.Errorf("Merge(m1, m2) = %v, want %v", got, want)
	}
}

func TestMergeInPlace(t *testing.T) {
	m1 := map[int]string{1: "one", 2: "two"}
	m2 := map[int]string{2: "deux", 3: "trois"}
	MergeInPlace(m1, m2)
	want := map[int]string{1: "one", 2: "deux", 3: "trois"}
	if !maps.Equal(m1, want) {
		t.Errorf("MergeInPlace(m1, m2) = %v, want %v", m1, want)
	}
}
