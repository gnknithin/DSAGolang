package linkedlist

import (
	"testing"
)

func TestIterateList(t *testing.T) {
	cases := []struct {
		in   []int
		want string
	}{
		{{1}, ""},
	}
	for _, c := range cases {
		got = Iterate(c.in)
	}
}
