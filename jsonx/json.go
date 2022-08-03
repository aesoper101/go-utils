package jsonx

import "encoding/json"

func MustJSON2String(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

func JSON2String(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}

func String2JSON(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}
