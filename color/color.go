package color

import (
	"fmt"
	"github.com/mgutz/ansi"
)

// colorize is a helper function that takes a color and a format string, then returns the colorized output.
func colorize(color string, format string, a ...interface{}) string {
	colorCode := ansi.ColorCode(color)
	resetCode := ansi.ColorCode("reset")
	return fmt.Sprintf("%s%s%s", colorCode, fmt.Sprintf(format, a...), resetCode)
}

// Green is a shorthand function for colorizing text in green
func Green(format string, a ...interface{}) string {
	return colorize("green", format, a...)
}

// Red is a shorthand function for colorizing text in red
func Red(format string, a ...interface{}) string {
	return colorize("red", format, a...)
}

// Yellow is a shorthand function for colorizing text in yellow
func Yellow(format string, a ...interface{}) string {
	return colorize("yellow+b+h", format, a...)
}
