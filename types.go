package cloudloyalty_client

import (
	"strconv"
	"time"
)

type IntOrString int

func (i *IntOrString) UnmarshalJSON(b []byte) error {
	if len(b) >= 2 && b[0] == '"' && b[len(b)-1] == '"' {
		// empty string is considered as zero
		if len(b) == 2 {
			*i = IntOrString(0)
			return nil
		}
		b = b[1 : len(b)-1]
	}
	val, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		val = 0
	}
	*i = IntOrString(val)
	return nil
}

type Birthdate time.Time

// UnmarshalJSON parses string to date. Following formats are supported:
// - RFC3999
// - YYYY-MM-DD
// - MM-DD
func (d *Birthdate) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	// note: the value is wrapped with "
	t, err := time.Parse("\"01-02\"", string(b))
	if err != nil {
		t, err = time.Parse("\"2006-01-02\"", string(b))
		if err != nil {
			t, err = time.Parse(`"`+time.RFC3339+`"`, string(b))
			if err != nil {
				return err
			}
		}
	}
	y := t.Year()
	if y < 1900 || y > time.Now().Year()+1 { // Allow to set currentYear + 1 for expecting a baby
		y = 1900
	}
	t = time.Date(y, t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	*d = Birthdate(t)
	return nil
}

func (d *Birthdate) MarshalJSON() ([]byte, error) {
	return time.Time(*d).MarshalJSON()
}
