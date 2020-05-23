package cloudloyalty_client

import (
	"testing"
	"time"
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
		var i IntOrString
		if err := i.UnmarshalJSON([]byte(c.input)); err != nil || int(i) != c.expected {
			t.Fatalf("failed asserting that \"%s\" produces %d (got %d)", c.input, c.expected, i)
		}
	}
}

func TestBirthdate_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		input    string
		expected string
		err      string
	}{
		{
			`"1986-09-07T00:00:00+03:00"`,
			"1986-09-07T00:00:00Z",
			"",
		},
		{
			`"1986-09-07T15:59:59-01:00"`,
			"1986-09-07T00:00:00Z",
			"",
		},
		{
			`"1986-09-07"`,
			"1986-09-07T00:00:00Z",
			"",
		},
		{
			`"09-07"`,
			"1900-09-07T00:00:00Z",
			"",
		},
		{
			`"1800-09-07"`,
			"1900-09-07T00:00:00Z",
			"",
		},
		{
			`"2800-09-07"`,
			"1900-09-07T00:00:00Z",
			"",
		},
		{
			`"1986-09-77"`,
			"1900-09-07T00:00:00Z",
			`parsing time ""1986-09-77"" as ""2006-01-02T15:04:05Z07:00"": cannot parse """ as "T"`,
		},
		{
			`""`, // empty string
			"1900-09-07T00:00:00Z",
			`parsing time """" as ""2006-01-02T15:04:05Z07:00"": cannot parse """ as "2006"`,
		},
		{
			"null", // JSON null
			"0001-01-01T00:00:00Z",
			"",
		},
	}

	for _, c := range cases {
		var d Birthdate
		err := d.UnmarshalJSON([]byte(c.input))
		if err != nil {
			if err.Error() != c.err {
				t.Fatalf("failed asserting that \"%s\" parse error is matched \"%s\" (got \"%s\")", c.input, c.err, err)
			}
			continue
		}
		expected, _ := time.Parse(time.RFC3339, c.expected)
		if time.Time(d) != expected {
			t.Fatalf("failed asserting that \"%s\" produces %s (got %s)", c.input, c.expected, time.Time(d))
		}
	}
}
