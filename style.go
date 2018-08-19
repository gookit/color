package color

import (
	"fmt"
)

// Style a colored style
// can add: fg color, bg color, color options
// quick use:
// 	color.Style(color.FgGreen).
type Style []Color

// New create a custom style
func New(colors ...Color) Style {
	return Style(colors)
}

// Save save to styles map
func (s Style) Save(name string) {
	AddStyle(name, s)
}

// Render render text
// usage:
//  color.New(color.FgGreen).Render("text")
//  color.New(color.FgGreen, color.BgBlack, color.OpBold).Render("text")
func (s Style) Render(args ...interface{}) string {
	if isLikeInCmd {
		return fmt.Sprint(args...)
	}

	return buildColoredText(buildColorCode(s...), args...)
}

// Sprint is alias of the 'Render'
func (s Style) Sprint(args ...interface{}) string {
	return s.Render(args...)
}

// Print render and Print text
func (s Style) Print(args ...interface{}) (int, error) {
	if isLikeInCmd {
		return winPrint(fmt.Sprint(args...), s...)
	}

	return fmt.Print(s.Render(args...))
}

// Printf render and Print text
func (s Style) Printf(format string, args ...interface{}) (int, error) {
	str := fmt.Sprintf(format, args...)
	if isLikeInCmd {
		return winPrint(str, s...)
	}

	return fmt.Print(s.Render(str))
}

// Println render and print text line
func (s Style) Println(args ...interface{}) (int, error) {
	if isLikeInCmd {
		return winPrintln(fmt.Sprint(args...), s...)
	}

	return fmt.Println(s.Render(args...))
}

// IsEmpty style
func (s Style) IsEmpty() bool {
	return len(s) == 0
}

/*************************************************************
 * internal styles(like bootstrap style)
 *************************************************************/

// internal styles(like bootstrap style)
// usage:
//	color.Info.Print("message")
//	color.Info.Println("new line")
//	color.Info.Printf("a %s message", "test")
//	color.Warn.Println("message")
//	color.Error.Println("message")
var (
	// Info color style
	Info = Style{OpReset, FgGreen}
	// Note color style
	Note = Style{OpBold, FgLightCyan}
	// Warn color style
	Warn = Style{OpBold, FgYellow}
	// Light color style
	Light = Style{FgLightWhite}
	// Error color style
	Error = Style{FgLightWhite, BgRed}
	// Danger color style
	Danger = Style{OpBold, FgRed}
	// Notice color style
	Notice = Style{OpBold, FgCyan}
	// Comment color style
	Comment = Style{OpReset, FgMagenta}
	// Success color style
	Success = Style{OpBold, FgGreen}
	// Primary color style
	Primary = Style{OpReset, FgBlue}
	// Question color style
	Question = Style{OpReset, FgMagenta}
	// Secondary color style
	Secondary = Style{FgDarkGray}
)

// Some defined styles, like bootstrap styles
var Styles = map[string]Style{
	"info":  Info,
	"note":  Note,
	"light": Light,
	"error": Error,

	"danger":  Danger,
	"notice":  Notice,
	"success": Success,
	"comment": Comment,
	"primary": Primary,
	"warning": Warn,

	"question":  Question,
	"secondary": Secondary,
}

// some style name alias
var styleAliases = map[string]string{
	"err":  "error",
	"suc":  "success",
	"warn": "warning",
}

// AddStyle add a style
func AddStyle(name string, s Style) {
	Styles[name] = s
}

// GetStyle get style by name
func GetStyle(name string) Style {
	if s, ok := Styles[name]; ok {
		return s
	}

	if realName, ok := styleAliases[name]; ok {
		return Styles[realName]
	}

	// empty style
	return New()
}

// GetStyleName
func GetStyleName(name string) string {
	if realName, ok := styleAliases[name]; ok {
		return realName
	}

	return name
}
