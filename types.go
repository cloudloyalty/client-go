package cloudloyalty_client

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type IntAsIntOrString int

// UnmarshalJSON decodes int or string into int without triggering an error.
// If string is not a valid number, assumes 0.
func (i *IntAsIntOrString) UnmarshalJSON(b []byte) error {
	if len(b) >= 2 && b[0] == '"' && b[len(b)-1] == '"' {
		// empty string is considered as zero
		if len(b) == 2 {
			*i = IntAsIntOrString(0)
			return nil
		}
		b = b[1 : len(b)-1]
	}
	val, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		val = 0
	}
	*i = IntAsIntOrString(val)
	return nil
}

type Birthdate time.Time
type ChildBirthdate time.Time

// UnmarshalJSON parses string to date. Following formats are supported:
// - RFC3999
// - YYYY-MM-DD
// - MM-DD
// If the actual year is unknown, 1900 should be passed.
func (d *Birthdate) UnmarshalJSON(b []byte) error {
	t, err := parseBirthdate(b, false)
	if err != nil {
		return err
	}
	if t != nil {
		*d = Birthdate(*t)
	}
	return nil
}

func (d *ChildBirthdate) UnmarshalJSON(b []byte) error {
	t, err := parseBirthdate(b, true)
	if err != nil {
		return err
	}
	if t != nil {
		*d = ChildBirthdate(*t)
	}
	return nil
}

// MarshalJSON formats birthdate as a JSON string. Date format is YYYY-MM-DD.
// Year 1900 means that the actual year is unknown.
func (d *Birthdate) MarshalJSON() ([]byte, error) {
	t := time.Time(*d)
	if t.Year() < 1900 || t.Year() > time.Now().Year()+1 {
		t = time.Date(1900, t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	}
	b := make([]byte, 0, 12)
	b = append(b, '"')
	b = t.AppendFormat(b, "2006-01-02")
	b = append(b, '"')
	return b, nil
}

func (d *ChildBirthdate) MarshalJSON() ([]byte, error) {
	b := Birthdate(*d)
	bRef := &b
	return bRef.MarshalJSON()
}

func parseBirthdate(b []byte, child bool) (*time.Time, error) {
	if string(b) == "null" {
		return nil, nil
	}
	// note: the value is wrapped with "
	t, err := time.Parse("\"01-02\"", string(b))
	if err != nil {
		t, err = time.Parse("\"2006-01-02\"", string(b))
		if err != nil {
			t, err = time.Parse(`"`+time.RFC3339+`"`, string(b))
			if err != nil {
				return nil, err
			}
		}
	}
	y := t.Year()
	if y < 1900 ||
		y > time.Now().Year()+1 || // For children allow to set currentYear + 1 for expecting a baby
		(!child && y >= time.Now().Year()-1) { // For adults the year must be less than currentYear - 1
		y = 1900
	}
	t = time.Date(y, t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	return &t, nil
}

type ExtraFields map[string]interface{}

type IntOrAuto struct {
	json.Unmarshaler

	Auto  bool
	Value int
}

func (i *IntOrAuto) UnmarshalJSON(v []byte) error {
	if strings.EqualFold(string(v), "\"auto\"") {
		i.Auto = true
		i.Value = 0
		return nil
	}
	i.Auto = false
	iv, err := strconv.Atoi(string(v))
	if err != nil {
		return err
	}
	i.Value = iv
	return nil
}

func (i *IntOrAuto) MarshalJSON() ([]byte, error) {
	if i.Auto {
		return []byte("\"auto\""), nil
	}
	return []byte(strconv.Itoa(i.Value)), nil
}
