package gocolors

import (
	"fmt"
	"strings"
)

// 8 Bit colors
const (
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

// 16 Bit colors
const (
	fgBrightBlack   = "\u001b[30;1m"
	fgBrightRed     = "\u001b[31;1m"
	fgBrightGreen   = "\u001b[32;1m"
	fgBrightYellow  = "\u001b[33;1m"
	fgBrightBlue    = "\u001b[34;1m"
	fgBrightMagenta = "\u001b[35;1m"
	fgBrightCyan    = "\u001b[36;1m"
	fgBrightWhite   = "\u001b[37;1m"
)

// Decorations
const (
	bold      = "\u001b[1m"
	underline = "\u001b[4m"
	reversed  = "\u001b[7m"
)

// Color returns a string rendered with the given color.
func Color(content, color, decoration string) string {
	switch strings.ToLower(color) {
	case "black":
		return Black(content, decoration)
	case "red":
		return Red(content, decoration)
	case "green":
		return Green(content, decoration)
	case "yellow":
		return Yellow(content, decoration)
	case "blue":
		return Blue(content, decoration)
	case "magenta":
		return Magenta(content, decoration)
	case "cyan":
		return Cyan(content, decoration)
	case "white":
		return White(content, decoration)
	case "bright-black":
		return BrightBlack(content, decoration)
	case "bright-red":
		return BrightRed(content, decoration)
	case "bright-green":
		return BrightGreen(content, decoration)
	case "bright-yellow":
		return BrightYellow(content, decoration)
	case "bright-blue":
		return BrightBlue(content, decoration)
	case "bright-magenta":
		return BrightMagenta(content, decoration)
	case "bright-cyan":
		return BrightCyan(content, decoration)
	case "bright-white":
		return BrightWhite(content, decoration)
	}

	return ""
}

// Black returns a string rendered in black.
func Black(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fmt.Sprint(fgBlack, bold, content, Reset)
	case "underline":
		return fmt.Sprint(fgBlack, underline, content, Reset)
	case "reversed":
		return fmt.Sprint(fgBlack, reversed, content, Reset)
	default:
		return fmt.Sprint(fgBlack, content, Reset)
	}
}

// BrightBlack returns a string rendered in bright black.
func BrightBlack(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fmt.Sprint(fgBrightBlack, bold, content, Reset)
	case "underline":
		return fmt.Sprint(fgBrightBlack, underline, content, Reset)
	case "reversed":
		return fmt.Sprint(fgBrightBlack, reversed, content, Reset)
	default:
		return fmt.Sprint(fgBrightBlack, content, Reset)
	}
}

// Red returns a string rendered in red.
func Red(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fmt.Sprint(fgRed, bold, content, Reset)
	case "underline":
		return fmt.Sprint(fgRed, underline, content, Reset)
	case "reversed":
		return fmt.Sprint(fgRed, reversed, content, Reset)
	default:
		return fmt.Sprint(fgRed, content, Reset)
	}
}

// BrightRed returns a string rendered in bright red.
func BrightRed(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fmt.Sprint(fgBrightRed, bold, content, Reset)
	case "underline":
		return fmt.Sprint(fgBrightRed, underline, content, Reset)
	case "reversed":
		return fmt.Sprint(fgBrightRed, reversed, content, Reset)
	default:
		return fmt.Sprint(fgBrightRed, content, Reset)
	}
}

// Green returns a string rendered in green.
func Green(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fmt.Sprint(fgGreen, bold, content, Reset)
	case "underline":
		return fmt.Sprint(fgGreen, underline, content, Reset)
	case "reversed":
		return fmt.Sprint(fgGreen, reversed, content, Reset)
	default:
		return fmt.Sprint(fgGreen, content, Reset)
	}
}

// BrightGreen returns a string rendered in bright green.
func BrightGreen(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fmt.Sprint(fgBrightGreen, bold, content, Reset)
	case "underline":
		return fmt.Sprint(fgBrightGreen, underline, content, Reset)
	case "reversed":
		return fmt.Sprint(fgBrightGreen, reversed, content, Reset)
	default:
		return fmt.Sprint(fgBrightGreen, content, Reset)
	}
}

// Yellow returns a string rendered in yellow.
func Yellow(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fmt.Sprint(fgYellow, bold, content, Reset)
	case "underline":
		return fmt.Sprint(fgYellow, underline, content, Reset)
	case "reversed":
		return fmt.Sprint(fgYellow, reversed, content, Reset)
	default:
		return fmt.Sprint(fgYellow, content, Reset)
	}
}

// BrightYellow returns a string rendered in bright yellow.
func BrightYellow(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fmt.Sprint(fgBrightYellow, bold, content, Reset)
	case "underline":
		return fmt.Sprint(fgBrightYellow, underline, content, Reset)
	case "reversed":
		return fmt.Sprint(fgBrightYellow, reversed, content, Reset)
	default:
		return fmt.Sprint(fgBrightYellow, content, Reset)
	}
}

// Blue returns a string rendered in blue.
func Blue(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fmt.Sprint(fgBlue, bold, content, Reset)
	case "underline":
		return fmt.Sprint(fgBlue, underline, content, Reset)
	case "reversed":
		return fmt.Sprint(fgBlue, reversed, content, Reset)
	default:
		return fmt.Sprint(fgBlue, content, Reset)
	}
}

// BrightBlue returns a string rendered in bright blue.
func BrightBlue(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fmt.Sprint(fgBrightBlue, bold, content, Reset)
	case "underline":
		return fmt.Sprint(fgBrightBlue, underline, content, Reset)
	case "reversed":
		return fmt.Sprint(fgBrightBlue, reversed, content, Reset)
	default:
		return fmt.Sprint(fgBrightBlue, content, Reset)
	}
}

// Magenta returns a string rendered in magenta.
func Magenta(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fmt.Sprint(fgMagenta, bold, content, Reset)
	case "underline":
		return fmt.Sprint(fgMagenta, underline, content, Reset)
	case "reversed":
		return fmt.Sprint(fgMagenta, reversed, content, Reset)
	default:
		return fmt.Sprint(fgMagenta, content, Reset)
	}
}

// BrightMagenta returns a string rendered in bright magenta.
func BrightMagenta(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fmt.Sprint(fgBrightMagenta, bold, content, Reset)
	case "underline":
		return fmt.Sprint(fgBrightMagenta, underline, content, Reset)
	case "reversed":
		return fmt.Sprint(fgBrightMagenta, reversed, content, Reset)
	default:
		return fmt.Sprint(fgBrightMagenta, content, Reset)
	}
}

// Cyan returns a string rendered in cyan.
func Cyan(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fmt.Sprint(fgCyan, bold, content, Reset)
	case "underline":
		return fmt.Sprint(fgCyan, underline, content, Reset)
	case "reversed":
		return fmt.Sprint(fgCyan, reversed, content, Reset)
	default:
		return fmt.Sprint(fgCyan, content, Reset)
	}
}

// BrightCyan returns a string rendered in bright cyan.
func BrightCyan(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fmt.Sprint(fgBrightCyan, bold, content, Reset)
	case "underline":
		return fmt.Sprint(fgBrightCyan, underline, content, Reset)
	case "reversed":
		return fmt.Sprint(fgBrightCyan, reversed, content, Reset)
	default:
		return fmt.Sprint(fgBrightCyan, content, Reset)
	}
}

// White returns a string rendered in white.
func White(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fmt.Sprint(fgWhite, bold, content, Reset)
	case "underline":
		return fmt.Sprint(fgWhite, underline, content, Reset)
	case "reversed":
		return fmt.Sprint(fgWhite, reversed, content, Reset)
	default:
		return fmt.Sprint(fgWhite, content, Reset)
	}
}

// BrightWhite returns a string rendered in bright white.
func BrightWhite(content, decoration string) string {
	switch strings.ToLower(decoration) {
	case "bold":
		return fmt.Sprint(fgBrightWhite, bold, content, Reset)
	case "underline":
		return fmt.Sprint(fgBrightWhite, underline, content, Reset)
	case "reversed":
		return fmt.Sprint(fgBrightWhite, reversed, content, Reset)
	default:
		return fmt.Sprint(fgBrightWhite, content, Reset)
	}
}
