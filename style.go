package color

import (
	"fmt"
	"strings"
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

// Printf render and print text
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
 * Theme
 *************************************************************/

// Theme definition. extends from Style
type Theme struct {
	// Name theme name
	Name string
	// Style for the theme
	Style
}

func NewTheme(name string, style Style) *Theme {
	return &Theme{name, style}
}

// Tips use name as title, only apply style for name
func (t *Theme) Tips(format string, args ...interface{}) {
	title := strings.ToUpper(t.Name) + ": "
	t.Print(title) // only apply style for name
	Printf(format+"\n", args...)
}

// Prompt use name as title, and apply style for message
func (t *Theme) Prompt(format string, args ...interface{}) {
	title := strings.ToUpper(t.Name) + ": "
	t.Printf(title+format+"\n", args...)
}

// Block like Prompt, but will wrap a empty line
func (t *Theme) Block(format string, args ...interface{}) {
	title := strings.ToUpper(t.Name) + ":\n  "
	message := Sprintf(format, args...)

	t.Println(title, message)
}

/*************************************************************
 * internal themes
 *************************************************************/

// internal themes(like bootstrap style)
// usage:
// 	color.Info.Print("message")
// 	color.Info.Println("new line")
// 	color.Info.Printf("a %s message", "test")
// 	color.Warn.Println("message")
// 	color.Error.Println("message")
var (
	// Info color style
	Info = &Theme{"info", Style{OpReset, FgGreen}}
	// Note color style
	Note = &Theme{"note", Style{OpBold, FgLightCyan}}
	// Warn color style
	Warn = &Theme{"warning", Style{OpBold, FgYellow}}
	// Light color style
	Light = &Theme{"light", Style{FgLightWhite}}
	// Error color style
	Error = &Theme{"error", Style{FgLightWhite, BgRed}}
	// Danger color style
	Danger = &Theme{"danger", Style{OpBold, FgRed}}
	// Notice color style
	Notice = &Theme{"notice", Style{OpBold, FgCyan}}
	// Comment color style
	Comment = &Theme{"comment", Style{OpReset, FgLightYellow}}
	// Success color style
	Success = &Theme{"success", Style{OpBold, FgGreen}}
	// Primary color style
	Primary = &Theme{"primary", Style{OpReset, FgBlue}}
	// Question color style
	Question = &Theme{"question", Style{OpReset, FgMagenta}}
	// Secondary color style
	Secondary = &Theme{"secondary", Style{FgDarkGray}}
)

// Themes internal defined themes.
// usage:
// 	color.Themes["info"].Println("message")
var Themes = map[string]*Theme{
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

// AddTheme add a theme and style
func AddTheme(name string, style Style) {
	Themes[name] = NewTheme(name, style)
	Styles[name] = style
}

// GetTheme get defined theme by name
func GetTheme(name string) *Theme {
	return Themes[name]
}

/*************************************************************
 * internal styles
 *************************************************************/

// Styles internal defined styles, like bootstrap styles.
// usage:
// 	color.Styles["info"].Println("message")
var Styles = map[string]Style{
	"info":  {OpReset, FgGreen},
	"note":  {OpBold, FgLightCyan},
	"light": {FgLightWhite, BgRed},
	"error": {FgLightWhite, BgRed},

	"danger":  {OpBold, FgRed},
	"notice":  {OpBold, FgCyan},
	"success": {OpBold, FgGreen},
	"comment": {OpReset, FgMagenta},
	"primary": {OpReset, FgBlue},
	"warning": {OpBold, FgYellow},

	"question":  {OpReset, FgMagenta},
	"secondary": {FgDarkGray},
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

// GetStyle get defined style by name
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
