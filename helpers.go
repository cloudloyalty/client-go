package cloudloyalty_client

import "time"

func FormatDateTime(t time.Time) string {
	res := t.Format(time.RFC3339)
	if res[len(res)-1:] == "Z" {
		res = res[:len(res)-1] + "+00:00"
	}
	return res
}

func ParseDateTime(in string) (*time.Time, error) {
	res, err := time.Parse(time.RFC3339, in)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func StringPtr(a string) *string    { return &a }
func IntPtr(a int) *int             { return &a }
func Float64Ptr(a float64) *float64 { return &a }
func BoolPtr(a bool) *bool          { return &a }
