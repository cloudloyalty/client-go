package cloudloyalty_client

import (
	"encoding/json"
	"strconv"
	"strings"
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
