package cloudloyalty_client

import (
	"testing"
)

func TestIntOrString_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{
			"null",
			0,
		},
		{
			"0",
			0,
		},
		{
			"10",
			10,
		},
		{
			"10.467",
			10,
		},
		{
			"\"non-int\"",
			0,
		},
		{
			"",
			0,
		},
		{
			"-10",
			-10,
		},
		{
			"-10.445",
			-10,
		},
	}

	for _, c := range cases {
		var i IntAsIntOrString
		if err := i.UnmarshalJSON([]byte(c.input)); err != nil || int(i) != c.expected {
			t.Fatalf("failed asserting that \"%s\" produces %d (got %d)", c.input, c.expected, i)
		}
	}
}
