package color

import (
	"fmt"
	"strings"
)

// Color represents a text color.
type Color uint8

// reset color
const Reset Color = 0

// Foreground colors.
const (
	// basic Foreground colors 30 - 37
	FgBlack   Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite

	FgDefault Color = 39

	// extra Foreground color 90 - 97
	FgDarkGray     Color = iota + 90
	FgLightRed
	FgLightGreen
	FgLightYellow
	FgLightBlue
	FgLightMagenta
	FgLightCyan
	FgWhiteEx
)

// Foreground colors map
var FgColors = map[string]Color{
	"black":   FgBlack,
	"red":     FgRed,
	"green":   FgGreen,
	"yellow":  FgYellow,
	"blue":    FgBlue,
	"magenta": FgMagenta,
	"cyan":    FgCyan,
	"white":   FgWhite,
	"default": FgDefault,
}

// Background colors.
const (
	// basic Background colors 40 - 47
	BgBlack   Color = iota + 40
	BgRed
	BgGreen
	BgYellow   // BgBrown like yellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	BgDefault Color = 49

	// extra Background color 100 - 107
	BgDarkGray     Color = iota + 100
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgWhiteEx
)

// Background colors map
var BgColors = map[string]Color{
	"black":   BgBlack,
	"red":     BgRed,
	"green":   BgGreen,
	"yellow":  BgYellow,
	"blue":    BgBlue,
	"magenta": BgMagenta,
	"cyan":    BgCyan,
	"white":   BgWhite,
	"default": BgDefault,
}

// color options
const (
	OpBold       = 1 // 加粗
	OpFuzzy      = 2 // 模糊(不是所有的终端仿真器都支持)
	OpItalic     = 3 // 斜体(不是所有的终端仿真器都支持)
	OpUnderscore = 4 // 下划线
	OpBlink      = 5 // 闪烁
	OpReverse    = 7 // 颠倒的 交换背景色与前景色
	OpConcealed  = 8 // 隐匿的
)

// color options map
var Options = map[string]Color{
	"bold":       OpBold,
	"fuzzy":      OpFuzzy,
	"italic":     OpItalic,
	"underscore": OpUnderscore,
	"blink":      OpBlink,
	"reverse":    OpReverse,
	"concealed":  OpConcealed,
}

// CLI color template
// const FullColorTpl = "\033[%sm%s\033[0m"
const FullColorTpl = "\x1b[%sm%s\x1b[0m"
const SingleColorTpl = "\x1b[%dm%s\x1b[0m"

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

// String to string
func (c Color) String() string {
	return fmt.Sprintf("%d", uint8(c))
}

// IsFgColor
func IsFgColor(name string) bool {
	if _, ok := FgColors[name]; ok {
		return true
	}

	return false
}

// IsBgColor
func IsBgColor(name string) bool {
	if _, ok := BgColors[name]; ok {
		return true
	}

	return false
}

// IsOption
func IsOption(name string) bool {
	if _, ok := Options[name]; ok {
		return true
	}

	return false
}

// buildColorCode return like "32;45;3"
func buildColorCode(colors ...Color) string {
	var codes []string

	for _, color := range colors {
		codes = append(codes, color.String())
	}

	return strings.Join(codes, ";")
}

// buildColoredText
func buildColoredText(code string, str string) string {
	return fmt.Sprintf(FullColorTpl, code, str)
}
