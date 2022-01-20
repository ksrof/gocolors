package gocolors

import (
	"strings"
)

// 8 Bit colors
var (
	fgBlack   = "\u001b[30m"
	fgRed     = "\u001b[31m"
	fgGreen   = "\u001b[32m"
	fgYellow  = "\u001b[33m"
	fgBlue    = "\u001b[34m"
	fgMagenta = "\u001b[35m"
	fgCyan    = "\u001b[36m"
	fgWhite   = "\u001b[37m"
	Reset     = "\u001b[0m"
)

// Decorations
var (
	bold      = "\u001b[1m"
	underline = "\u001b[4m"
	reversed  = "\u001b[7m"
)

// Black renders a string with an ANSI black color code.
func Black(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fgBlack + bold + content + Reset
	case "underline":
		return fgBlack + underline + content + Reset
	case "reversed":
		return fgBlack + reversed + content + Reset
	default:
		return fgBlack + content + Reset
	}
}

// Red renders a string with an ANSI red color code.
func Red(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fgRed + bold + content + Reset
	case "underline":
		return fgRed + underline + content + Reset
	case "reversed":
		return fgRed + reversed + content + Reset
	default:
		return fgRed + content + Reset
	}
}

// Green renders a string with an ANSI green color code.
func Green(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fgGreen + bold + content + Reset
	case "underline":
		return fgGreen + underline + content + Reset
	case "reversed":
		return fgGreen + reversed + content + Reset
	default:
		return fgGreen + content + Reset
	}
}

// Yellow renders a string with an ANSI yellow color code.
func Yellow(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fgYellow + bold + content + Reset
	case "underline":
		return fgYellow + underline + content + Reset
	case "reversed":
		return fgYellow + reversed + content + Reset
	default:
		return fgYellow + content + Reset
	}
}

// Blue renders a string with an ANSI blue color code.
func Blue(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fgBlue + bold + content + Reset
	case "underline":
		return fgBlue + underline + content + Reset
	case "reversed":
		return fgBlue + reversed + content + Reset
	default:
		return fgBlue + content + Reset
	}
}

// Magenta renders a string with an ANSI magenta color code.
func Magenta(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fgMagenta + bold + content + Reset
	case "underline":
		return fgMagenta + underline + content + Reset
	case "reversed":
		return fgMagenta + reversed + content + Reset
	default:
		return fgMagenta + content + Reset
	}
}

// Cyan renders a string with an ANSI cyan color code.
func Cyan(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fgCyan + bold + content + Reset
	case "underline":
		return fgCyan + underline + content + Reset
	case "reversed":
		return fgCyan + reversed + content + Reset
	default:
		return fgCyan + content + Reset
	}
}

// White renders a string with an ANSI white color code.
func White(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fgWhite + bold + content + Reset
	case "underline":
		return fgWhite + underline + content + Reset
	case "reversed":
		return fgWhite + reversed + content + Reset
	default:
		return fgWhite + content + Reset
	}
}
