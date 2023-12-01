package fp

import (
	"reflect"
	"testing"
)

func TestMapEmptyReturnsEmpty(t *testing.T) {
	in := []string{}
	mapper := func(t string) string {
		return ""
	}
	want := []string{}

	got := Map(in, mapper)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestStringToInt(t *testing.T) {
	in := []string{"lorem", "ipsum", "dolor", "sit", "amet"}
	mapper := func(t string) int {
		return len(t)
	}
	want := []int{5, 5, 5, 3, 4}

	got := Map(in, mapper)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
