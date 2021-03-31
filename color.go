/*
Package color is Command line color library.
Support rich color rendering output, universal API method, compatible with Windows system

Source code and other details for the project are available at GitHub:

	https://github.com/gookit/color

More usage please see README and tests.
*/
package color

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/xo/terminfo"
)

// color render templates
// ESC 操作的表示:
// 	"\033"(Octal 8进制) = "\x1b"(Hexadecimal 16进制) = 27 (10进制)
const (
	SettingTpl   = "\x1b[%sm"
	FullColorTpl = "\x1b[%sm%s\x1b[0m"
)

// ResetSet Close all properties.
const ResetSet = "\x1b[0m"

// CodeExpr regex to clear color codes eg "\033[1;36mText\x1b[0m"
const CodeExpr = `\033\[[\d;?]+m`

var (
	// Enable switch color render and display
	Enable = true
	// RenderTag render HTML tag on call color.Xprint, color.PrintX
	RenderTag = true
	// errors on windows render OR print to io.Writer
	errors []error
	// output the default io.Writer message print
	output io.Writer = os.Stdout
	// mark current env, It's like in `cmd.exe`
	// if not in windows, it's always is False.
	isLikeInCmd bool
	// match color codes
	codeRegex = regexp.MustCompile(CodeExpr)
	// mark current env is support color.
	// Always: isLikeInCmd != supportColor
	// supportColor = IsSupportColor()
)

/*************************************************************
 * global settings
 *************************************************************/

// Set set console color attributes
func Set(colors ...Color) (int, error) {
	if !Enable { // not enable
		return 0, nil
	}

	if !SupportColor() {
		return 0, nil
	}

	return fmt.Printf(SettingTpl, Colors2code(colors...))
}

// Reset reset console color attributes
func Reset() (int, error) {
	if !Enable { // not enable
		return 0, nil
	}

	if !SupportColor() {
		return 0, nil
	}

	return fmt.Print(ResetSet)
}

// Disable disable color output
func Disable() bool {
	oldVal := Enable
	Enable = false
	return oldVal
}

// NotRenderTag on call color.Xprint, color.PrintX
func NotRenderTag() {
	RenderTag = false
}

// SetOutput set default colored text output
func SetOutput(w io.Writer) {
	output = w
}

// ResetOutput reset output
func ResetOutput() {
	output = os.Stdout
}

// ResetOptions reset all package option setting
func ResetOptions() {
	RenderTag = true
	Enable = true
	output = os.Stdout
}

// SupportColor on the current ENV
func SupportColor() bool {
	return colorLevel > terminfo.ColorLevelNone
}

// ForceColor force open color render
func ForceSetColorLevel(level terminfo.ColorLevel) terminfo.ColorLevel {
	oldLevelVal := colorLevel
	colorLevel = level
	return oldLevelVal
}

// ForceColor force open color render
func ForceColor() terminfo.ColorLevel {
	return ForceOpenColor()
}

// ForceOpenColor force open color render
func ForceOpenColor() terminfo.ColorLevel {
	// TODO should set level to ?
	return ForceSetColorLevel(terminfo.ColorLevelMillions)
}

// IsLikeInCmd check result
func IsLikeInCmd() bool {
	return isLikeInCmd
}

// GetErrors info
func GetErrors() []error {
	return errors
}

/*************************************************************
 * render color code
 *************************************************************/

// RenderCode render message by color code.
// Usage:
// 	msg := RenderCode("3;32;45", "some", "message")
func RenderCode(code string, args ...interface{}) string {
	var message string
	if ln := len(args); ln == 0 {
		return ""
	}

	message = fmt.Sprint(args...)
	if len(code) == 0 {
		return message
	}

	// disabled OR not support color
	if !Enable || !SupportColor() {
		return ClearCode(message)
	}

	return fmt.Sprintf(FullColorTpl, code, message)
}

// RenderWithSpaces Render code with spaces.
// If the number of args is > 1, a space will be added between the args
func RenderWithSpaces(code string, args ...interface{}) string {
	message := formatArgsForPrintln(args)
	if len(code) == 0 {
		return message
	}

	// disabled OR not support color
	if !Enable || !SupportColor() {
		return ClearCode(message)
	}

	return fmt.Sprintf(FullColorTpl, code, message)
}

// RenderString render a string with color code.
// Usage:
// 	msg := RenderString("3;32;45", "a message")
func RenderString(code string, str string) string {
	if len(code) == 0 || str == "" {
		return str
	}

	// disabled OR not support color
	if !Enable || !SupportColor() {
		return ClearCode(str)
	}

	return fmt.Sprintf(FullColorTpl, code, str)
}

// ClearCode clear color codes.
// eg:
// 		"\033[36;1mText\x1b[0m" -> "Text"
func ClearCode(str string) string {
	return codeRegex.ReplaceAllString(str, "")
}
