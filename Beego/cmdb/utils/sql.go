package utils

import (
	"fmt"
	"strings"
)

func Like(q string) string {
	if q != "" {
		strings.TrimSpace(q)
		q = strings.ReplaceAll(q, "/", "//")
		q = strings.ReplaceAll(q, "%", "/%")
		q = strings.ReplaceAll(q, "_", "/_")
		return fmt.Sprintf("%%%s%%", q)
	}
	return ""
}
