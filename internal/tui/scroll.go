package tui

import (
	"strings"

	"github.com/charmbracelet/x/ansi"
)

func applyHorizontalScroll(s string, scrollX, width int) string {
	if scrollX <= 0 {
		return ansi.Truncate(s, width, "")
	}

	stripped := ansi.Strip(s)
	if ansi.StringWidth(stripped) <= scrollX {
		return ""
	}

	cutPlain := cutLeftPlain(stripped, scrollX)
	truncated := ansi.Truncate(cutPlain, width, "")

	return truncated
}

func cutLeftPlain(s string, width int) string {
	if width <= 0 {
		return s
	}
	w := 0
	for i := 0; i < len(s); {
		r, size := decodeRune(s, i)
		rw := ansi.StringWidth(string(r))
		if w+rw > width {
			return s[i:]
		}
		w += rw
		i += size
	}
	return ""
}

func decodeRune(s string, i int) (rune, int) {
	r := rune(s[i])
	if r < 0x80 {
		return r, 1
	}
	size := 1
	if r&0xE0 == 0xC0 {
		size = 2
	} else if r&0xF0 == 0xE0 {
		size = 3
	} else if r&0xF8 == 0xF0 {
		size = 4
	}
	if i+size > len(s) {
		size = len(s) - i
	}
	var result rune
	for j := 0; j < size; j++ {
		result = (result << 8) | rune(s[i+j])
	}
	return result, size
}

func scrollLines(s string, scrollX, width int) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = applyHorizontalScroll(line, scrollX, width)
	}
	return strings.Join(lines, "\n")
}

func wrapContent(content string, width int) string {
	if width <= 0 {
		return content
	}

	wrapped := ansi.Wordwrap(content, width, " ")

	lines := strings.Split(wrapped, "\n")
	for i, line := range lines {
		stripped := ansi.Strip(line)
		if ansi.StringWidth(stripped) > width {
			lines[i] = ansi.Truncate(line, width, "")
		}
	}

	return strings.Join(lines, "\n")
}
