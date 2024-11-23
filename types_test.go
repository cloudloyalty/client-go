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

func TestIntOrAuto_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		input        string
		expectedAuto bool
		expectedInt  int
		expectedErr  bool
	}{
		{
			"null",
			false,
			0,
			true,
		},
		{
			"0",
			false,
			0,
			false,
		},
		{
			"10",
			false,
			10,
			false,
		},
		{
			"10.467",
			false,
			0,
			true,
		},
		{
			"\"non-int\"",
			false,
			0,
			true,
		},
		{
			"",
			false,
			0,
			true,
		},
		{
			"-10",
			false,
			-10,
			false,
		},
		{
			"-10.445",
			false,
			0,
			true,
		},
		{
			"-10.0",
			false,
			-10,
			false,
		},
		{
			"\"auto\"",
			true,
			0,
			false,
		},
		{
			"\"AUTO\"",
			true,
			0,
			false,
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			var i IntOrAuto
			err := i.UnmarshalJSON([]byte(c.input))
			if c.expectedErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !c.expectedErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if i.Auto != c.expectedAuto {
				t.Fatalf("expected Auto to be %t, got %t", c.expectedAuto, i.Auto)
			}
			if int(i.Value) != c.expectedInt {
				t.Fatalf("expected Value to be %d, got %d", c.expectedInt, i.Value)
			}
		})
	}
}
