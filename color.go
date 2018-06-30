package color

import (
	"fmt"
	"strings"
	"regexp"
	"github.com/gookit/cliapp/utils"
)

// Color represents a text color.
type Color uint8

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

	// revert default FG
	FgDefault Color = 39

	// extra Foreground color 90 - 97(非标准)
	FgDarkGray     Color = iota + 90 // 亮黑（灰）
	FgLightRed
	FgLightGreen
	FgLightYellow
	FgLightBlue
	FgLightMagenta
	FgLightCyan
	FgWhiteEx
)

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

	// revert default BG
	BgDefault Color = 49

	// extra Background color 100 - 107(非标准)
	BgDarkGray     Color = iota + 100
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgWhiteEx
)

// color options
const (
	OpReset      = 0 // 重置所有设置
	OpBold       = 1 // 加粗
	OpFuzzy      = 2 // 模糊(不是所有的终端仿真器都支持)
	OpItalic     = 3 // 斜体(不是所有的终端仿真器都支持)
	OpUnderscore = 4 // 下划线
	OpBlink      = 5 // 闪烁
	OpReverse    = 7 // 颠倒的 交换背景色与前景色
	OpConcealed  = 8 // 隐匿的
)

// ESC 操作的表示 "\033"(Octal 8进制) = "\x1b"(Hexadecimal 16进制) = 27 (10进制)
const ResetCode = "\x1b[0m"

// CLI color template
const SettingTpl = "\x1b[%sm"
const FullColorTpl = "\x1b[%sm%s\x1b[0m"
const SingleColorTpl = "\x1b[%dm%s\x1b[0m"

// Regex to clear color codes eg "\033[1;36mText\x1b[0m"
const CodeExpr = `\033\[[\d;?]+m`

// switch color display
var Enable = true
var isSupportColor = utils.IsSupportColor()

// init
func init() {
	// Byte8Color("test 8 byte color", 0x98)
	// os.Exit(0)
}

// Set
func Set(colors ...Color) (int, error) {
	return fmt.Printf(SettingTpl, buildColorCode(colors...))
}

// SetByCode
func SetByCode(code string) (int, error) {
	return fmt.Printf(SettingTpl, code)
}

// Reset
func Reset() (int, error) {
	return fmt.Print(ResetCode)
}

// Disable disable color output
func Disable() {
	Enable = false
}

// Render
func (c Color) Render(args ...interface{}) string {
	str := fmt.Sprint(args...)

	return fmt.Sprintf(SingleColorTpl, uint8(c), str)
}

// Renderf
func (c Color) Renderf(format string, args ...interface{}) string {
	str := fmt.Sprintf(format, args...)

	return fmt.Sprintf(SingleColorTpl, uint8(c), str)
}

// Print
func (c Color) Print(args ...interface{}) (int, error) {
	return fmt.Print(c.Render(args...))
}

// Println
func (c Color) Println(args ...interface{}) (int, error) {
	return fmt.Println(c.Render(args...))
}

// Printf
// usage:
// 	color.FgCyan.Printf("string %s", "arg0")
func (c Color) Printf(format string, args ...interface{}) (int, error) {
	return fmt.Print(c.Renderf(format, args...))
}

// IsValid 检测是否为一个有效的 Color 值
func (c Color) IsValid() bool {
	return c < 107
}

// String to string
func (c Color) String() string {
	return fmt.Sprintf("%d", uint8(c))
}

// Apply apply custom colors
// usage:
// 	// (string, fg-color,bg-color, options...)
//  color.Apply("text", color.FgGreen)
//  color.Apply("text", color.FgGreen, color.BgBlack, color.OpBold)
func Apply(str string, colors ...Color) string {
	return buildColoredText(
		buildColorCode(colors...),
		str,
	)
}

// RenderCodes "3;32;45"
func RenderCodes(code string, str string) string {
	return buildColoredText(code, str)
}

// ClearCode clear color codes
// eg "\033[36;1mText\x1b[0m" -> "Text"
func ClearCode(str string) string {
	reg := regexp.MustCompile(CodeExpr)
	// r1 := reg.FindAllString("\033[36;1mText\x1b[0m", -1)

	return reg.ReplaceAllString(str, "")
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
func buildColoredText(code string, args ...interface{}) string {
	str := fmt.Sprint(args...)

	if len(code) == 0 {
		return str
	}

	if !Enable {
		return ClearCode(str)
	}

	// if not support color output
	if !isSupportColor {
		return ClearCode(str)
	}

	return fmt.Sprintf(FullColorTpl, code, str)
}

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

// color options map
var Options = map[string]Color{
	"reset":      OpReset,
	"bold":       OpBold,
	"fuzzy":      OpFuzzy,
	"italic":     OpItalic,
	"underscore": OpUnderscore,
	"blink":      OpBlink,
	"reverse":    OpReverse,
	"concealed":  OpConcealed,
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
