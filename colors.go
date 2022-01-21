package gocolors

import (
	"fmt"
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
		render := fmt.Sprint(fgBlack, bold, content, Reset)
		return render
	case "underline":
		render := fmt.Sprint(fgBlack, underline, content, Reset)
		return render
	case "reversed":
		render := fmt.Sprint(fgBlack, reversed, content, Reset)
		return render
	default:
		render := fmt.Sprint(fgBlack, content, Reset)
		return render
	}
}

// Red renders a string with an ANSI red color code.
func Red(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		render := fmt.Sprint(fgRed, bold, content, Reset)
		return render
	case "underline":
		render := fmt.Sprint(fgRed, underline, content, Reset)
		return render
	case "reversed":
		render := fmt.Sprint(fgRed, reversed, content, Reset)
		return render
	default:
		render := fmt.Sprint(fgRed, content, Reset)
		return render
	}
}

// Green renders a string with an ANSI green color code.
func Green(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		render := fmt.Sprint(fgGreen, bold, content, Reset)
		return render
	case "underline":
		render := fmt.Sprint(fgGreen, underline, content, Reset)
		return render
	case "reversed":
		render := fmt.Sprint(fgGreen, reversed, content, Reset)
		return render
	default:
		render := fmt.Sprint(fgGreen, content, Reset)
		return render
	}
}

// Yellow renders a string with an ANSI yellow color code.
func Yellow(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		render := fmt.Sprint(fgYellow, bold, content, Reset)
		return render
	case "underline":
		render := fmt.Sprint(fgYellow, underline, content, Reset)
		return render
	case "reversed":
		render := fmt.Sprint(fgYellow, reversed, content, Reset)
		return render
	default:
		render := fmt.Sprint(fgYellow, content, Reset)
		return render
	}
}

// Blue renders a string with an ANSI blue color code.
func Blue(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		render := fmt.Sprint(fgBlue, bold, content, Reset)
		return render
	case "underline":
		render := fmt.Sprint(fgBlue, underline, content, Reset)
		return render
	case "reversed":
		render := fmt.Sprint(fgBlue, reversed, content, Reset)
		return render
	default:
		render := fmt.Sprint(fgBlue, content, Reset)
		return render
	}
}

// Magenta renders a string with an ANSI magenta color code.
func Magenta(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		render := fmt.Sprint(fgMagenta, bold, content, Reset)
		return render
	case "underline":
		render := fmt.Sprint(fgMagenta, underline, content, Reset)
		return render
	case "reversed":
		render := fmt.Sprint(fgMagenta, reversed, content, Reset)
		return render
	default:
		render := fmt.Sprint(fgMagenta, content, Reset)
		return render
	}
}

// Cyan renders a string with an ANSI cyan color code.
func Cyan(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		render := fmt.Sprint(fgCyan, bold, content, Reset)
		return render
	case "underline":
		render := fmt.Sprint(fgCyan, underline, content, Reset)
		return render
	case "reversed":
		render := fmt.Sprint(fgCyan, reversed, content, Reset)
		return render
	default:
		render := fmt.Sprint(fgCyan, content, Reset)
		return render
	}
}

// White renders a string with an ANSI white color code.
func White(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		render := fmt.Sprint(fgWhite, bold, content, Reset)
		return render
	case "underline":
		render := fmt.Sprint(fgWhite, underline, content, Reset)
		return render
	case "reversed":
		render := fmt.Sprint(fgWhite, reversed, content, Reset)
		return render
	default:
		render := fmt.Sprint(fgWhite, content, Reset)
		return render
	}
}
