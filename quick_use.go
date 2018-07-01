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

/**
	print and add new line
 */

// Boldln add bold for message
func Boldln(args ...interface{}) {
	OpBold.Println(args...)
}

// Blackln add black for message
func Blackln(args ...interface{}) {
	FgBlack.Println(args...)
}

// Whiteln add white for message
func Whiteln(args ...interface{})  {
	FgWhite.Println(args...)
}

// Grayln add gray for message
func Grayln(args ...interface{}) {
	FgDarkGray.Println(args...)
}

// Redln add red for message
func Redln(args ...interface{}) {
	FgRed.Println(args...)
}

// Greenln add green for message
func Greenln(args ...interface{}) {
	FgGreen.Println(args...)
}

// Yellowln add yellow for message
func Yellowln(args ...interface{}) {
	FgYellow.Println(args...)
}

// Blueln add blue for message
func Blueln(args ...interface{}) {
	FgBlue.Println(args...)
}

// Magentaln add magenta for message
func Magentaln(args ...interface{}) {
	FgMagenta.Println(args...)
}

// Cyanln add cyan for message
func Cyanln(args ...interface{}) {
	FgCyan.Println(args...)
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

/**
	print and add new line
 */

// Suc add suc style for message
func Sucln(args ...interface{}) {
	GetStyle("success").Println(args...)
}

// Success add success style for message
func Successln(args ...interface{}) {
	GetStyle("success").Println(args...)
}

// Info add info style for message
func Infoln(args ...interface{}) {
	GetStyle("info").Println(args...)
}

// Secondary add secondary style for message
func Secondaryln(args ...interface{}) {
	GetStyle("secondary").Println(args...)
}

// Comment add comment style for message
func Commentln(args ...interface{}) {
	GetStyle("comment").Println(args...)
}

// Note add note style for message
func Noteln(args ...interface{}) {
	GetStyle("note").Println(args...)
}

// Notice add notice style for message
func Noticeln(args ...interface{}) {
	GetStyle("notice").Println(args...)
}

// Light add light style for message
func Lightln(args ...interface{}) {
	GetStyle("light").Println(args...)
}

// Warn add warn style for message
func Warnln(args ...interface{}) {
	GetStyle("warning").Println(args...)
}

// Warning add warning style for message
func Warningln(args ...interface{}) {
	GetStyle("warning").Println(args...)
}

// Primary add primary style for message
func Primaryln(args ...interface{}) {
	GetStyle("primary").Println(args...)
}

// Danger add danger style for message
func Dangerln(args ...interface{}) {
	GetStyle("danger").Println(args...)
}

// Err add err style for message
func Errln(args ...interface{}) {
	GetStyle("error").Println(args...)
}

// Error add error style for message
func Errorln(args ...interface{}) {
	GetStyle("error").Println(args...)
}

// Question add question style for message
func Questionln(args ...interface{}) {
	GetStyle("question").Println(args...)
}
