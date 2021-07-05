package mdhtml

import (
	"errors"
	"fmt"
	"strings"
)

var heading = "<h%d>%s</h%d>"

func makeHeader(level int, line string) (string, error) {
	if level < 1 || level > 6 {
		return "", fmt.Errorf("Unsupported nesting level: %d", level)
	}

	return fmt.Sprintf(heading, level, strings.TrimSpace(line[level+1:]), level), nil
}

func Trunslate(text []string) (string, error) {
	builder := strings.Builder{}
	for _, line := range text {
		nestingLevel := 1
		switch {
		case strings.HasPrefix(line, "###### "):
			nestingLevel += 1
			fallthrough
		case strings.HasPrefix(line, "##### "):
			nestingLevel += 1
			fallthrough
		case strings.HasPrefix(line, "#### "):
			nestingLevel += 1
			fallthrough
		case strings.HasPrefix(line, "### "):
			nestingLevel += 1
			fallthrough
		case strings.HasPrefix(line, "## "):
			nestingLevel += 1
			fallthrough
		case strings.HasPrefix(line, "# "):
			newLine, err := makeHeader(nestingLevel, line)
			if err != nil {
				return "", err
			}
			_, err = builder.WriteString(newLine)
			if err != nil {
				return "", err
			}
		default:
			return "", errors.New("not implemented yet")
		}
	}
	return builder.String(), nil
}
