package cloudloyalty_client

import (
	"encoding/json"
	"fmt"
	"math"
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
	// it is allowed for value to be a float with zero fractional part, e.g. 1.0
	f, err := strconv.ParseFloat(string(v), 64)
	if err != nil {
		return err
	}
	iv, frac := math.Modf(f)
	if frac != 0 {
		return fmt.Errorf("unexpected fractional part in integer value: %f", f)
	}
	i.Value = int(iv)
	return nil
}

func (i *IntOrAuto) MarshalJSON() ([]byte, error) {
	if i.Auto {
		return []byte("\"auto\""), nil
	}
	return []byte(strconv.Itoa(i.Value)), nil
}

type ValidRangeTime time.Time

var (
	TimeValidRangeStart = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	TimeValidRangeEnd   = time.Date(2040, 1, 1, 0, 0, 0, 0, time.UTC).Add(-time.Second)
)

func (v *ValidRangeTime) UnmarshalJSON(b []byte) error {
	t := (*time.Time)(v)
	if err := t.UnmarshalJSON(b); err != nil {
		return err
	}
	if t.Before(TimeValidRangeStart) || t.After(TimeValidRangeEnd) {
		return fmt.Errorf(
			"date %s is out of valid range (%s to %s)",
			t.Format(time.RFC3339),
			TimeValidRangeStart.Format(time.RFC3339),
			TimeValidRangeEnd.Format(time.RFC3339),
		)
	}
	return nil
}

func (v *ValidRangeTime) MarshalJSON() ([]byte, error) {
	t := (*time.Time)(v)
	return t.MarshalJSON()
}
