package color

// There are some funcs for quick start uss.

// type ColoredString string
//
// func (s ColoredString) String() string {
// 	return string(s)
// }
//
// func (s ColoredString) Print() {
// 	fmt.Print(s.String())
// }
//
// func (s ColoredString) Println() {
// 	fmt.Println(s.String())
// }
//
// // Bold use bold
// func Bold(args ...interface{}) ColoredString {
// 	return ColoredString(OpBold.Render(args...))
// }

// Bold add bold for message
func Bold(args ...interface{}) {
	OpBold.Print(args...)
}

// Black add black for message
func Black(args ...interface{}) {
	FgBlack.Print(args...)
}

// White add white for message
func White(args ...interface{})  {
	FgWhite.Print(args...)
}

// Gray add gray for message
func Gray(args ...interface{}) {
	FgDarkGray.Print(args...)
}

// Red add red for message
func Red(args ...interface{}) {
	FgRed.Print(args...)
}

// Green add green for message
func Green(args ...interface{}) {
	FgGreen.Print(args...)
}

// Yellow add yellow for message
func Yellow(args ...interface{}) {
	FgYellow.Print(args...)
}

// Blue add blue for message
func Blue(args ...interface{}) {
	FgBlue.Print(args...)
}

// Magenta add magenta for message
func Magenta(args ...interface{}) {
	FgMagenta.Print(args...)
}

// Cyan add cyan for message
func Cyan(args ...interface{}) {
	FgCyan.Print(args...)
}

// Some defined styles, like bootstrap styles
var BuiltinStyles = map[string]Style{
	"info":      {OpReset, FgGreen},
	"note":      {OpBold, FgLightCyan},
	"light":     {FgLightWhite},
	"error":     {FgLightWhite, BgRed},
	"danger":    {OpBold, FgRed},
	"notice":    {OpBold, FgCyan},
	"success":   {OpBold, FgGreen},
	"comment":   {OpReset, FgYellow},
	"primary":   {OpReset, FgBlue},
	"warning":   {OpBold, FgYellow},
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
	BuiltinStyles[name] = s
}

// GetStyle get style by name
func GetStyle(name string) Style {
	if s, ok := BuiltinStyles[name]; ok {
		return s
	}

	if realName, ok := styleAliases[name]; ok {
		return BuiltinStyles[realName]
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

// Suc add suc style for message
func Suc(args ...interface{}) {
	GetStyle("success").Print(args...)
}

// Success add success style for message
func Success(args ...interface{}) {
	GetStyle("success").Print(args...)
}

// Info add info style for message
func Info(args ...interface{}) {
	GetStyle("info").Print(args...)
}

// Secondary add secondary style for message
func Secondary(args ...interface{}) {
	GetStyle("secondary").Print(args...)
}

// Comment add comment style for message
func Comment(args ...interface{}) {
	GetStyle("comment").Print(args...)
}

// Note add note style for message
func Note(args ...interface{}) {
	GetStyle("note").Print(args...)
}

// Notice add notice style for message
func Notice(args ...interface{}) {
	GetStyle("notice").Print(args...)
}

// Light add light style for message
func Light(args ...interface{}) {
	GetStyle("light").Print(args...)
}

// Warn add warn style for message
func Warn(args ...interface{}) {
	GetStyle("warning").Print(args...)
}

// Warning add warning style for message
func Warning(args ...interface{}) {
	GetStyle("warning").Print(args...)
}

// Primary add primary style for message
func Primary(args ...interface{}) {
	GetStyle("primary").Print(args...)
}

// Danger add danger style for message
func Danger(args ...interface{}) {
	GetStyle("danger").Print(args...)
}

// Err add err style for message
func Err(args ...interface{}) {
	GetStyle("error").Print(args...)
}

// Error add error style for message
func Error(args ...interface{}) {
	GetStyle("error").Print(args...)
}

// Question add question style for message
func Question(args ...interface{}) {
	GetStyle("question").Print(args...)
}
