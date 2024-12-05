package htmlparser

import (
	"testing"

	"github.com/anaskhan96/soup"
)

func TestIsExampleAnnouncedTrue(t *testing.T) {
	cases := []soup.Root{
		soup.HTMLParse("<p>For example</p>").Find("p"),
		soup.HTMLParse("<p>For example, suppose you have the following list of contents from six rucksacks:</p>").Find("p"),
		soup.HTMLParse("<p>Here is an example engine schematic:</p>").Find("p"),
	}
	for _, tc := range cases {
		t.Run(tc.HTML(), func(t *testing.T) {
			if !IsExampleAnnounced(tc) {
				t.Errorf("IsExampleAnnounced(%v) = false, want true", tc)
			}
		})
	}
}

func TestIsExampleAnnouncedFalse(t *testing.T) {
	cases := []soup.Root{
		soup.HTMLParse("<p>So, in the above example, the first group's rucksacks are the first three lines:</p>").Find("p"),
	}
	for _, tc := range cases {
		t.Run(tc.HTML(), func(t *testing.T) {
			if IsExampleAnnounced(tc) {
				t.Errorf("IsExampleAnnounced(%v) = true, want false", tc)
			}
		})
	}
}
