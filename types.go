package cloudloyalty_client

import (
	"strconv"
)

type IntOrString int

func (i *IntOrString) UnmarshalJSON(b []byte) error {
	if len(b) >= 3 && b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	val, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}
	*i = IntOrString(val)
	return nil
}
