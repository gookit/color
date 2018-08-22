package color

import (
	"fmt"
	"regexp"
	"strings"
)

// Color represents a text color. 3/4 bite color.
// ESC 操作的表示:
// 	"\033"(Octal 8进制) = "\x1b"(Hexadecimal 16进制) = 27 (10进制)
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

// color render templates
const (
	SettingTpl       = "\x1b[%sm"
	FullColorTpl     = "\x1b[%sm%s\x1b[0m"
	FullColorNlTpl   = "\x1b[%sm%s\x1b[0m\n"
	SingleColorTpl   = "\x1b[%dm%s\x1b[0m"
	SingleColorNlTpl = "\x1b[%dm%s\x1b[0m\n"
)

// ResetCode value
const ResetCode = "0"

// ResetSet 重置/正常 关闭所有属性。
const ResetSet = "\x1b[0m"

// CodeExpr regex to clear color codes eg "\033[1;36mText\x1b[0m"
const CodeExpr = `\033\[[\d;?]+m`

// Enable switch color display
var Enable = true

var (
	// mark current env, It's like in cmd.exe
	isLikeInCmd bool
	// match color codes
	codeRegex = regexp.MustCompile(CodeExpr)
	// check current env
	isSupportColor = IsSupportColor()
)

/*************************************************************
 * global settings
 *************************************************************/

// Set set console color attributes
func Set(colors ...Color) (int, error) {
	if !Enable { // not enable
		return 0, nil
	}

	// on windows cmd.exe
	if isLikeInCmd {
		return winSet(colors...)
	}

	return fmt.Printf(SettingTpl, colors2code(colors...))
}

// Reset reset console color attributes
func Reset() (int, error) {
	if !Enable { // not enable
		return 0, nil
	}

	// on windows cmd.exe
	if isLikeInCmd {
		return winReset()
	}

	return fmt.Print(ResetSet)
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

// Text render a text message
func (c Color) Text(message string) string {
	if isLikeInCmd {
		return message
	}

	return fmt.Sprintf(SingleColorTpl, c, message)
}

// Render messages by color setting
// usage:
// 		green := color.FgGreen.Render
// 		fmt.Println(green("message"))
func (c Color) Render(a ...interface{}) string {
	message := fmt.Sprint(a...)
	if isLikeInCmd {
		return message
	}

	return fmt.Sprintf(SingleColorTpl, c, message)
}

// Sprint render messages by color setting. is alias of the Render()
func (c Color) Sprint(a ...interface{}) string {
	return c.Render(a...)
}

// Sprintf format and render message.
// Usage:
// 	green := color.FgGreen.RenderFn()
//  colored := green("message")
func (c Color) Sprintf(format string, args ...interface{}) string {
	message := fmt.Sprintf(format, args...)
	if isLikeInCmd {
		return message
	}

	return fmt.Sprintf(SingleColorTpl, c, message)
}

// Print messages.
// Usage:
// 		color.Green.Print("message")
// OR:
// 		green := color.FgGreen.Print
// 		green("message")
func (c Color) Print(args ...interface{}) (int, error) {
	if isLikeInCmd {
		return winPrint(fmt.Sprint(args...), c)
	}

	return fmt.Printf(SingleColorTpl, c, fmt.Sprint(args...))
}

// Printf format and print messages.
// usage:
// 		color.Cyan.Printf("string %s", "arg0")
func (c Color) Printf(format string, args ...interface{}) (int, error) {
	if isLikeInCmd {
		return winPrint(fmt.Sprintf(format, args...), c)
	}

	return fmt.Printf(SingleColorTpl, c, fmt.Sprintf(format, args...))
}

// Println messages with new line
func (c Color) Println(args ...interface{}) (int, error) {
	if isLikeInCmd {
		return winPrintln(fmt.Sprint(args...), c)
	}

	return fmt.Printf(SingleColorNlTpl, c, fmt.Sprint(args...))
}

// String to code string. eg "35"
func (c Color) String() string {
	return fmt.Sprintf("%d", c)
}

// IsValid color value
func (c Color) IsValid() bool {
	return c < 107
}

/*************************************************************
 * render color code
 *************************************************************/

// RenderCode render message by color code.
// Usage:
// 	msg := RenderCode("3;32;45", "some", "message")
func RenderCode(code string, args ...interface{}) string {
	message := fmt.Sprint(args...)
	if len(code) == 0 || isLikeInCmd {
		return message
	}

	if !Enable {
		return ClearCode(message)
	}

	// if not support color output
	if !isSupportColor {
		return ClearCode(message)
	}

	return fmt.Sprintf(FullColorTpl, code, message)
}

// RenderString render a string with color code.
// Usage:
// 	msg := RenderString("3;32;45", "a message")
func RenderString(code string, message string) string {
	// some check
	if isLikeInCmd || len(code) == 0 || message == "" {
		return message
	}

	if !Enable {
		return ClearCode(message)
	}

	// if not support color output
	if !isSupportColor {
		return ClearCode(message)
	}

	return fmt.Sprintf(FullColorTpl, code, message)
}

/*************************************************************
 * helper methods
 *************************************************************/

// Apply custom colors.
// Usage:
// 		// (string, fg-color,bg-color, options...)
//  	color.Apply("text", color.FgGreen)
//  	color.Apply("text", color.FgGreen, color.BgBlack, color.OpBold)
func Apply(message string, colors ...Color) string {
	return RenderCode(colors2code(colors...), message)
}

// ClearCode clear color codes.
// eg:
// 		"\033[36;1mText\x1b[0m" -> "Text"
func ClearCode(str string) string {
	return codeRegex.ReplaceAllString(str, "")
}

// convert colors to code. return like "32;45;3"
func colors2code(colors ...Color) string {
	if len(colors) == 0 {
		return ""
	}

	var codes []string
	for _, color := range colors {
		codes = append(codes, color.String())
	}

	return strings.Join(codes, ";")
}
