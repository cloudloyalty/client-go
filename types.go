package cloudloyalty_client

import (
	"strconv"
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
