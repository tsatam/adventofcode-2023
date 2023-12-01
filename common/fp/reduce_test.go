package fp

import (
	"reflect"
	"testing"
)

func TestReduceEmptyReturnsIdentity(t *testing.T) {
	in := []int{}
	identity := 42
	combine := func(curr int, next int) int {
		return -12
	}

	want := identity
	got := Reduce(in, identity, combine)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReduceIntsSumReturnsSum(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	identity := 0
	combine := func(curr int, next int) int {
		return curr + next
	}

	want := 15
	got := Reduce(in, identity, combine)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
