package aoc

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestInts(t *testing.T) {
	cases := []struct {
		input string
		want  []int
	}{
		{
			input: "-1   100   42 ",
			want:  []int{-1, 100, 42},
		},
		{
			input: "-1,100,+42 ",
			want:  []int{-1, 100, 42},
		},
		{
			input: "????.######..#####. 1,6,5",
			want:  []int{1, 6, 5},
		},
	}
	for _, test := range cases {
		t.Run(test.input, func(t *testing.T) {
			got := Ints(test.input)
			if !slices.Equal(got, test.want) {
				t.Fatalf("Unexpected Int() result: got %v, want %v", got, test.want)
			}
		})
	}
}
