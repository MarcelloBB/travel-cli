package utils

import (
	"strings"
)

func ParseHeaders(headers string) map[string]string {
	parsed := make(map[string]string)

	// Divide por ponto e vírgula
	pairs := strings.Split(headers, ";")
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}
		parts := strings.SplitN(pair, ":", 2)
		if len(parts) != 2 {
			continue // ignora headers malformados
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		parsed[key] = value
	}

	return parsed
}
