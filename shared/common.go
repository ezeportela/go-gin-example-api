package shared

import (
	"encoding/json"
	"time"
)

func StringifyInterface(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func CastInterface(src interface{}, dst interface{}) {
	b, _ := json.Marshal(src)
	json.Unmarshal(b, &dst)
}

func FormatDateTime(time time.Time) string {
	return time.Format("2006-01-02T15:04:05.111")
}
