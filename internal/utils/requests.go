package utils

import (
	"bytes"
	"encoding/json"
	"strings"
)

func ParseHeaders(headers string) map[string]string {
	parsed := make(map[string]string)

	pairs := strings.Split(headers, ";")
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}
		parts := strings.SplitN(pair, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		parsed[key] = value
	}

	return parsed
}

func PrettyPrintJSON(raw string) (bytes.Buffer, error) {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(raw), "", "  ")
	if err != nil {
		return bytes.Buffer{}, err
	}
	return out, nil
}
