package color

import (
	"fmt"
	"strings"
)

// Color represents a text color.
type Color uint8

// CLI color template
// const FullColorTpl = "\033[%sm%s\033[0m"
const FullColorTpl = "\x1b[%sm%s\x1b[0m"
const SingleColorTpl = "\x1b[%dm%s\x1b[0m"

// New
func New(c Color) *Color {
	return &c
}

// S adds the coloring to the given string.
// usage: cli.Color(cli.FgCyan).S("string")
func (c Color) S(s string) string {
	return fmt.Sprintf(SingleColorTpl, uint8(c), s)
}

// F adds the coloring to the given string.
// usage: cli.Color(cli.FgCyan).F("string %s", "arg0")
func (c Color) F(s string, args ...interface{}) string {
	s = fmt.Sprintf(s, args...)

	return fmt.Sprintf(SingleColorTpl, uint8(c), s)
}

// Print
func (c Color) Print(args ...interface{}) (int, error) {
	str := fmt.Sprint(args...)

	return fmt.Printf(SingleColorTpl, uint8(c), str)
}

// Println
func (c Color) Println(args ...interface{}) (int, error) {
	str := fmt.Sprintln(args...)

	return fmt.Printf(SingleColorTpl, uint8(c), str)
}

// Printf
func (c Color) Printf(format string, args ...interface{}) (int, error) {
	str := fmt.Sprintf(format, args...)

	return fmt.Printf(SingleColorTpl, uint8(c), str)
}

// IsValid 检测是否为一个有效的 Color 值
func (c Color) IsValid() bool {
	return c < 107
}

// String to string
func (c Color) String() string {
	return fmt.Sprintf("%d", uint8(c))
}

// UseCodes "32;45;3"
func UseCodes(code string, str string) string {
	return buildColoredText(code, str)
}

// buildColorCode return like "32;45;3"
func buildColorCode(colors ...Color) string {
	if len(colors) == 0 {
		return ""
	}

	var codes []string

	for _, color := range colors {
		codes = append(codes, color.String())
	}

	return strings.Join(codes, ";")
}

// buildColoredText
func buildColoredText(code string, str ...string) string {
	if len(code) == 0 {
		return strings.Join(str, "")
	}

	return fmt.Sprintf(FullColorTpl, code, strings.Join(str, ""))
}
