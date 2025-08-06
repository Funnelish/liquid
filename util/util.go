package util

import (
	"fmt"
	"html"
)

func ErrorPlaceholder(line int, msg string) string {
	return fmt.Sprintf(
		`<span class="liquid-error hidden" data-line="%d">%s</span>`,
		line,
		html.EscapeString(msg),
	)
}
