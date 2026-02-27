package utils

import "encoding/json"

func ParseJSON(data any) ([]byte, error) {
	return json.Marshal(data)
}

func UnmarshalJSON(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
