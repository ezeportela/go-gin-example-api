package shared

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"
)

func StringifyInterface(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func ParseJSON(src string, dst interface{}) error {
	err := json.Unmarshal([]byte(src), dst)
	return err
}

func JsonToStruct(src interface{}, dst interface{}) error {
	attrs := StringifyInterface(src)
	return ParseJSON(attrs, &dst)
}

func CastInterface(src interface{}, dst interface{}) {
	b, _ := json.Marshal(src)
	json.Unmarshal(b, &dst)
}

func FormatDateTime(time time.Time) string {
	return time.Format("2006-01-02T15:04:05.111")
}

func ReadFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)

	return data, err
}

func MakeRoute(basePath string, path string) string {
	return strings.Join([]string{basePath, path}, "")
}
