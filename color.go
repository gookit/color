package color

import (
	"fmt"
	"regexp"
	"strings"
)

// Color represents a text color.
// 3/4 bite color.
type Color uint8

// Foreground colors. basic foreground colors 30 - 37
const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta // 品红
	FgCyan    // 青色
	FgWhite
	// FgDefault revert default FG
	FgDefault Color = 39
)

// Extra foreground color 90 - 97(非标准)
const (
	FgDarkGray Color = iota + 90 // 亮黑（灰）
	FgLightRed
	FgLightGreen
	FgLightYellow
	FgLightBlue
	FgLightMagenta
	FgLightCyan
	FgLightWhite
	// FgGray is alias of FgDarkGray
	FgGray Color = 90 // 亮黑（灰）
)

// Background colors. basic background colors 40 - 47
const (
	BgBlack Color = iota + 40
	BgRed
	BgGreen
	BgYellow // BgBrown like yellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	// BgDefault revert default BG
	BgDefault Color = 49
)

// Extra background color 100 - 107(非标准)
const (
	BgDarkGray Color = iota + 99
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgLightWhite
	// BgGray is alias of BgDarkGray
	BgGray Color = 100
)

// Option settings
const (
	OpReset         Color = iota // 0 重置所有设置
	OpBold                       // 1 加粗
	OpFuzzy                      // 2 模糊(不是所有的终端仿真器都支持)
	OpItalic                     // 3 斜体(不是所有的终端仿真器都支持)
	OpUnderscore                 // 4 下划线
	OpBlink                      // 5 闪烁
	OpFastBlink                  // 5 快速闪烁(未广泛支持)
	OpReverse                    // 7 颠倒的 交换背景色与前景色
	OpConcealed                  // 8 隐匿的
	OpStrikethrough              // 9 删除的，删除线(未广泛支持)
)

// console color mode
const (
	ModeNormal = iota
	Mode256    // 8 bite
	ModeRGB    // 24 bite
	ModeGrayscale
)

// There are basic foreground color alias
const (
	Red     = FgRed
	Cyan    = FgCyan
	Gray    = FgDarkGray
	Blue    = FgBlue
	Black   = FgBlack
	Green   = FgGreen
	White   = FgWhite
	Yellow  = FgYellow
	Magenta = FgMagenta
	Bold    = OpBold
	Normal  = FgDefault
)

// CLI color template
const (
	SettingTpl     = "\x1b[%sm"
	FullColorTpl   = "\x1b[%sm%s\x1b[0m"
	FullColorNlTpl   = "\x1b[%sm%s\x1b[0m\n"
	SingleColorTpl = "\x1b[%dm%s\x1b[0m"
)

const StartCode = "\x1b["

// ResetCode ESC 操作的表示
// 	"\033"(Octal 8进制) = "\x1b"(Hexadecimal 16进制) = 27 (10进制)
const ResetCode = "\x1b[0m"

// CodeExpr regex to clear color codes eg "\033[1;36mText\x1b[0m"
const CodeExpr = `\033\[[\d;?]+m`

// Enable switch color display
var Enable = true

// mark current env, It's like in cmd.exe
var isLikeInCmd bool

// check current env
var isSupportColor = IsSupportColor()

// Set set console color attributes
func Set(colors ...Color) (int, error) {
	// not enable
	if !Enable {
		return 0, nil
	}

	// on cmd.exe
	if isLikeInCmd {
		return winSet(colors...)
	}

	return fmt.Printf(SettingTpl, buildColorCode(colors...))
}

// Reset reset console color attributes
func Reset() (int, error) {
	// not enable
	if !Enable {
		return 0, nil
	}

	// on cmd.exe
	if isLikeInCmd {
		return winReset()
	}

	return fmt.Print(ResetCode)
}

// Disable disable color output
func Disable() {
	Enable = false
}

// IsDisabled color
func IsDisabled() bool {
	return Enable == false
}

/*************************************************************
 * basic render methods
 *************************************************************/

// Render messages by color setting
// usage:
// 		green := color.FgGreen.Render
// 		fmt.Println(green("message"))
func (c Color) Render(args ...interface{}) string {
	str := fmt.Sprint(args...)
	if isLikeInCmd {
		return str
	}

	return fmt.Sprintf(SingleColorTpl, c, str)
}

// Renderf format and render message
// usage:
// 	green := color.FgGreen.RenderFn()
//  colored := green("message")
func (c Color) Renderf(format string, args ...interface{}) string {
	str := fmt.Sprintf(format, args...)
	if isLikeInCmd {
		return str
	}

	return fmt.Sprintf(SingleColorTpl, c, str)
}

// Print messages
// usage:
// 		color.FgGreen.Print("message")
// or:
// 		green := color.FgGreen.Print
// 		green("message")
func (c Color) Print(args ...interface{}) (int, error) {
	if isLikeInCmd {
		return winPrint(fmt.Sprint(args...), c)
	}

	return fmt.Print(c.Render(args...))
}

// Println messages line
func (c Color) Println(args ...interface{}) (int, error) {
	if isLikeInCmd {
		return winPrintln(fmt.Sprint(args...), c)
	}

	return fmt.Println(c.Render(args...))
}

// Printf format and print messages
// usage:
// 		color.FgCyan.Printf("string %s", "arg0")
func (c Color) Printf(format string, args ...interface{}) (int, error) {
	if isLikeInCmd {
		return winPrint(fmt.Sprintf(format, args...), c)
	}

	return fmt.Print(c.Renderf(format, args...))
}

// IsValid 检测是否为一个有效的 Color 值
func (c Color) IsValid() bool {
	return c < 107
}

// String to string
func (c Color) String() string {
	return fmt.Sprintf("%d", c)
}

/*************************************************************
 * helper methods
 *************************************************************/

// Apply custom colors.
// usage:
// 		// (string, fg-color,bg-color, options...)
//  	color.Apply("text", color.FgGreen)
//  	color.Apply("text", color.FgGreen, color.BgBlack, color.OpBold)
func Apply(str string, colors ...Color) string {
	return buildColoredText(buildColorCode(colors...), str)
}

// RenderCodes render by color code "3;32;45"
func RenderCodes(code string, str string) string {
	return buildColoredText(code, str)
}

// ClearCode clear color codes
// eg:
// 		"\033[36;1mText\x1b[0m" -> "Text"
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
