package cloudloyalty_client

import "time"

func FormatDateTime(t time.Time) string {
	res := t.Format(time.RFC3339)
	if res[len(res)-1:] == "Z" {
		res = res[:len(res)-1] + "+00:00"
	}
	return res
}
